// kore-ai/proxy/circuitbreaker.go
package main

import (
	"sync"
	"time"
)

// CircuitState representa o estado atual do circuit breaker.
type CircuitState int

const (
	StateClosed   CircuitState = iota // operação normal, usa o modelo primário
	StateOpen                         // falhas excessivas, desvia todo o tráfego para o fallback
	StateHalfOpen                     // janela de teste após o cooldown
)

func (s CircuitState) String() string {
	switch s {
	case StateClosed:
		return "CLOSED"
	case StateOpen:
		return "OPEN"
	case StateHalfOpen:
		return "HALF_OPEN"
	default:
		return "UNKNOWN"
	}
}

// CircuitBreaker protege o modelo primário (ex: hermes3:70b) de sobrecarga, desviando
// tráfego para um modelo de fallback mais leve quando a taxa de erro estoura o threshold.
type CircuitBreaker struct {
	mu sync.Mutex

	primaryModel  string
	fallbackModel string

	failureThreshold float64       // taxa de erro (0.0 a 1.0) que dispara a abertura
	minRequests      int           // nº mínimo de requisições na janela antes de avaliar a taxa
	resetTimeout     time.Duration // tempo em OPEN antes de tentar HALF_OPEN

	state       CircuitState
	totalCount  int
	failCount   int
	lastFailure time.Time
	openedAt    time.Time
}

func NewCircuitBreaker(primary, fallback string, threshold float64, minRequests int, resetTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		primaryModel:     primary,
		fallbackModel:    fallback,
		failureThreshold: threshold,
		minRequests:      minRequests,
		resetTimeout:     resetTimeout,
		state:            StateClosed,
	}
}

// SelectModel decide, em tempo de requisição, qual modelo deve atender o pedido.
// Função de hot-path: O(1), trava o mutex pelo menor tempo possível.
func (cb *CircuitBreaker) SelectModel() string {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	switch cb.state {
	case StateOpen:
		if time.Since(cb.openedAt) >= cb.resetTimeout {
			cb.state = StateHalfOpen
			return cb.primaryModel // permite UMA tentativa de teste no modelo primário
		}
		return cb.fallbackModel
	case StateHalfOpen:
		return cb.primaryModel
	default: // StateClosed
		return cb.primaryModel
	}
}

// RecordResult deve ser chamado após cada chamada ao modelo primário para atualizar
// as estatísticas. success=false cobre timeouts e erros 5xx do Ollama.
func (cb *CircuitBreaker) RecordResult(success bool) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if cb.state == StateHalfOpen {
		if success {
			cb.resetLocked() // modelo primário se recuperou: fecha o circuito
			return
		}
		cb.openLocked() // falhou no teste: reabre e reinicia o cooldown
		return
	}

	cb.totalCount++
	if !success {
		cb.failCount++
		cb.lastFailure = time.Now()
	}

	if cb.totalCount >= cb.minRequests {
		rate := float64(cb.failCount) / float64(cb.totalCount)
		if rate >= cb.failureThreshold {
			cb.openLocked()
		}
	}
}

func (cb *CircuitBreaker) openLocked() {
	cb.state = StateOpen
	cb.openedAt = time.Now()
	cb.totalCount = 0
	cb.failCount = 0
}

func (cb *CircuitBreaker) resetLocked() {
	cb.state = StateClosed
	cb.totalCount = 0
	cb.failCount = 0
}

// CurrentState é exposto para fins de observabilidade (ex: endpoint /health).
func (cb *CircuitBreaker) CurrentState() CircuitState {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.state
}