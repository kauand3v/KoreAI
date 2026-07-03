// kore-ai/proxy/proxyconnector.go
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"kore-ai/core/types"
)

const ollamaChatEndpoint = "http://localhost:11434/api/chat"

// httpClient é reutilizado entre requisições (pool de conexões/keep-alive), evitando
// o custo de criar um client novo a cada chamada -- crítico para o SLA de baixa latência.
var httpClient = &http.Client{
	Timeout: 60 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     90 * time.Second,
	},
}

// ollamaMessage espelha o formato de mensagem esperado pela API nativa do Ollama.
type ollamaMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ollamaChatRequest é o payload traduzido do padrão OpenAI para o padrão Ollama.
type ollamaChatRequest struct {
	Model    string          `json:"model"`
	Messages []ollamaMessage `json:"messages"`
	Stream   bool            `json:"stream"` // explícito: false para resposta única, não-streamed
}

// ollamaChatResponse é a resposta crua devolvida pelo Ollama (modo não-streaming).
type ollamaChatResponse struct {
	Model           string        `json:"model"`
	CreatedAt       string        `json:"created_at"`
	Message         ollamaMessage `json:"message"`
	Done            bool          `json:"done"`
	PromptEvalCount int           `json:"prompt_eval_count"`
	EvalCount       int           `json:"eval_count"`
}

// ForwardToOllama traduz a requisição padrão OpenAI para o formato Ollama, executa a
// chamada HTTP local e devolve a resposta já normalizada de volta ao contrato OpenAI.
func ForwardToOllama(ctx context.Context, req types.OpenAIRequest, targetModel string) (types.OpenAIResponse, error) {
	ollamaMessages := make([]ollamaMessage, 0, len(req.Messages))
	for _, m := range req.Messages {
		ollamaMessages = append(ollamaMessages, ollamaMessage{Role: m.Role, Content: m.Content})
	}

	payload := ollamaChatRequest{
		Model:    targetModel,
		Messages: ollamaMessages,
		Stream:   false,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return types.OpenAIResponse{}, fmt.Errorf("falha ao serializar payload ollama: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, ollamaChatEndpoint, bytes.NewReader(body))
	if err != nil {
		return types.OpenAIResponse{}, fmt.Errorf("falha ao construir request ollama: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return types.OpenAIResponse{}, fmt.Errorf("falha de rede ao chamar ollama: %w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.OpenAIResponse{}, fmt.Errorf("falha ao ler resposta do ollama: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return types.OpenAIResponse{}, fmt.Errorf("ollama retornou status %d: %s", resp.StatusCode, string(respBytes))
	}

	var ollamaResp ollamaChatResponse
	if err := json.Unmarshal(respBytes, &ollamaResp); err != nil {
		return types.OpenAIResponse{}, fmt.Errorf("falha ao decodificar resposta ollama: %w", err)
	}

	return mapToOpenAIResponse(ollamaResp, targetModel), nil
}

// mapToOpenAIResponse converte o formato nativo do Ollama para o contrato OpenAI-compatível.
func mapToOpenAIResponse(o ollamaChatResponse, modelUsed string) types.OpenAIResponse {
	return types.OpenAIResponse{
		ID:      fmt.Sprintf("chatcmpl-%d", time.Now().UnixNano()),
		Object:  "chat.completion",
		Created: time.Now().Unix(),
		Model:   modelUsed,
		Choices: []types.Choice{
			{
				Index: 0,
				Message: types.OpenAIMessage{
					Role:    o.Message.Role,
					Content: o.Message.Content,
				},
				FinishReason: "stop",
			},
		},
		Usage: types.Usage{
			PromptTokens:     o.PromptEvalCount,
			CompletionTokens: o.EvalCount,
			TotalTokens:      o.PromptEvalCount + o.EvalCount,
		},
	}
}