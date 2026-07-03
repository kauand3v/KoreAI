// core/types/gateway-types.go
package types

import "time"

// OpenAIMessage representa uma única mensagem na conversa, no formato padrão OpenAI.
type OpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// OpenAIRequest é o contrato de entrada aceito pelo Edge Gateway -- compatível com
// o payload padrão da API /v1/chat/completions da OpenAI.
type OpenAIRequest struct {
	Model       string          `json:"model"`
	Messages    []OpenAIMessage `json:"messages"`
	Stream      bool            `json:"stream,omitempty"`
	Temperature float64         `json:"temperature,omitempty"`
	MaxTokens   int             `json:"max_tokens,omitempty"`
	// TenantID é uma extensão do KoreAI (fora do contrato oficial OpenAI), usada
	// para isolamento multi-tenant quando não enviada via header X-Tenant-ID.
	TenantID string `json:"tenant_id,omitempty"`
}

// Choice representa uma opção de resposta gerada pelo modelo.
type Choice struct {
	Index        int           `json:"index"`
	Message      OpenAIMessage `json:"message"`
	FinishReason string        `json:"finish_reason"`
}

// Usage contabiliza o consumo de tokens da requisição -- usado para billing e métricas.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// OpenAIResponse é o contrato de saída devolvido pelo Edge Gateway, normalizado a partir
// da resposta nativa do Ollama (ou de outros backends plugados no futuro).
type OpenAIResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

// GatewayConfig centraliza os parâmetros operacionais do gateway -- carregado uma vez
// no startup e compartilhado (somente leitura) entre todas as goroutines de requisição.
type GatewayConfig struct {
	OllamaURL               string
	PrimaryModel            string
	FallbackModel           string
	CacheTTL                time.Duration
	CircuitBreakerThreshold float64
	MaxFailures             int
	ResetTimeout            time.Duration
}

// MetricsPayload é a unidade de dado enviada via channel ao MetricsCollector --
// deve ser leve e sem ponteiros para evitar contenção de GC sob alta concorrência.
type MetricsPayload struct {
	TenantID   string    `json:"tenant_id"`
	Model      string    `json:"model"`
	LatencyMs  int64     `json:"latency_ms"`
	TokensUsed int       `json:"tokens_used"`
	CacheHit   bool      `json:"cache_hit"`
	Error      string    `json:"error,omitempty"`
	Timestamp  time.Time `json:"timestamp"`
}