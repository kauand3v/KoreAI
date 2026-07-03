// kore-ai/proxy/main.go
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"kore-ai/core/types"
)

// Estado global de orquestração -- inicializado uma única vez no startup e
// compartilhado (de forma thread-safe) entre todas as goroutines de requisição.
var (
	cache    SemanticCache
	breaker  *CircuitBreaker
	metrics  *MetricsCollector
	gwConfig types.GatewayConfig
)

func main() {
	gwConfig = types.GatewayConfig{
		OllamaURL:               "http://localhost:11434/api/chat",
		PrimaryModel:            "hermes3:70b",
		FallbackModel:           "hermes3:8b",
		CacheTTL:                10 * time.Minute,
		CircuitBreakerThreshold: 0.5, // abre acima de 50% de falhas na janela
		MaxFailures:             10,  // tamanho mínimo da janela de avaliação
		ResetTimeout:            30 * time.Second,
	}

	cache = NewRedisSemanticCache("localhost:6379", gwConfig.CacheTTL)
	breaker = NewCircuitBreaker(gwConfig.PrimaryModel, gwConfig.FallbackModel, gwConfig.CircuitBreakerThreshold, gwConfig.MaxFailures, gwConfig.ResetTimeout)
	metrics = NewMetricsCollector(1024, 4) // buffer de 1024 eventos, 4 workers consumidores

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/chat/completions", handleChatCompletions)
	mux.HandleFunc("/health", handleHealth)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Println("[kore-ai] edge gateway escutando na porta :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[kore-ai] erro fatal no servidor http: %v", err)
		}
	}()

	// Graceful shutdown: aguarda SIGINT/SIGTERM para drenar conexões e métricas pendentes.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("[kore-ai] sinal de shutdown recebido, drenando conexões...")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("[kore-ai] erro durante shutdown do servidor: %v", err)
	}
	metrics.Shutdown()
	log.Println("[kore-ai] shutdown completo.")
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"status":        "ok",
		"circuit_state": breaker.CurrentState().String(),
	})
}

// handleChatCompletions orquestra o fluxo completo de uma requisição no padrão OpenAI:
// 1) cache semântico -> 2) circuit breaker -> 3) proxy para o Ollama -> 4) métricas assíncronas.
func handleChatCompletions(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	if r.Method != http.MethodPost {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var req types.OpenAIRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "payload inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	tenantID := tenantFromRequest(r, req)
	prompt := lastUserPrompt(req)

	// 1) Cache semântico: hit retorna em microssegundos, sem tocar no modelo.
	if cached, hit := cache.CheckCache(tenantID, prompt); hit {
		resp := buildCachedResponse(cached, req.Model)
		writeJSON(w, http.StatusOK, resp)
		publishMetrics(tenantID, req.Model, start, resp.Usage.TotalTokens, true, "")
		return
	}

	// 2) Circuit breaker decide, em O(1), qual modelo deve receber o tráfego.
	targetModel := breaker.SelectModel()

	// 3) Encaminha para o Ollama com timeout dedicado para não travar o hot-path.
	ctx, cancel := context.WithTimeout(r.Context(), 45*time.Second)
	defer cancel()

	resp, err := ForwardToOllama(ctx, req, targetModel)
	if err != nil {
		breaker.RecordResult(false)
		publishMetrics(tenantID, targetModel, start, 0, false, err.Error())
		http.Error(w, "falha ao processar requisição no modelo: "+err.Error(), http.StatusBadGateway)
		return
	}
	breaker.RecordResult(true)

	// Persiste no cache de forma assíncrona -- nunca bloqueia a resposta ao cliente.
	go func(tid, p, content string) {
		if err := cache.SaveCache(tid, p, content); err != nil {
			log.Printf("[cache] falha ao persistir: %v", err)
		}
	}(tenantID, prompt, extractContent(resp))

	writeJSON(w, http.StatusOK, resp)
	publishMetrics(tenantID, targetModel, start, resp.Usage.TotalTokens, false, "")
}

// tenantFromRequest extrai o identificador do tenant via header dedicado, com
// fallback para o campo do body e depois para "default" em ambientes de teste.
func tenantFromRequest(r *http.Request, req types.OpenAIRequest) string {
	if t := r.Header.Get("X-Tenant-ID"); t != "" {
		return t
	}
	if req.TenantID != "" {
		return req.TenantID
	}
	return "default"
}

// lastUserPrompt extrai a última mensagem de role=user, usada como chave do cache semântico.
func lastUserPrompt(req types.OpenAIRequest) string {
	for i := len(req.Messages) - 1; i >= 0; i-- {
		if req.Messages[i].Role == "user" {
			return req.Messages[i].Content
		}
	}
	return ""
}

func extractContent(resp types.OpenAIResponse) string {
	if len(resp.Choices) == 0 {
		return ""
	}
	return resp.Choices[0].Message.Content
}

func buildCachedResponse(cachedContent, model string) types.OpenAIResponse {
	return types.OpenAIResponse{
		ID:      "chatcmpl-cache-hit",
		Object:  "chat.completion",
		Created: time.Now().Unix(),
		Model:   model,
		Choices: []types.Choice{
			{
				Index: 0,
				Message: types.OpenAIMessage{
					Role:    "assistant",
					Content: cachedContent,
				},
				FinishReason: "stop",
			},
		},
		Usage: types.Usage{}, // cache hit não consome tokens do modelo
	}
}

func publishMetrics(tenantID, model string, start time.Time, tokens int, cacheHit bool, errMsg string) {
	metrics.Publish(types.MetricsPayload{
		TenantID:   tenantID,
		Model:      model,
		LatencyMs:  time.Since(start).Milliseconds(),
		TokensUsed: tokens,
		CacheHit:   cacheHit,
		Error:      errMsg,
		Timestamp:  time.Now(),
	})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("[http] falha ao serializar resposta: %v", err)
	}
}