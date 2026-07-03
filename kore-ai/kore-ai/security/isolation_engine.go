package security

import (
	"sync"
	"time"
)

// TenantSecurityContext representa o estado de isolamento de um tenant,
// cacheado em memória na borda para checagens de baixíssima latência.
type TenantSecurityContext struct {
	TenantID    string
	Active      bool
	IsolationOK bool
	LastChecked time.Time
}

// IsolationEngine mantém o cache local de contextos de tenant.
// Em produção, o cache-miss/stale dispara uma revalidação contra o
// control plane (Python/Redis) — aqui simulado em memória.
type IsolationEngine struct {
	mu       sync.RWMutex
	tenants  map[string]*TenantSecurityContext
	cacheTTL time.Duration
}

func NewIsolationEngine(cacheTTL time.Duration) *IsolationEngine {
	return &IsolationEngine{
		tenants:  make(map[string]*TenantSecurityContext),
		cacheTTL: cacheTTL,
	}
}

// register simula a busca/validação do tenant no control plane e
// armazena o resultado no cache local.
func (e *IsolationEngine) register(tenantID string) *TenantSecurityContext {
	ctx := &TenantSecurityContext{
		TenantID:    tenantID,
		Active:      true, // TODO: substituir por chamada real ao control plane (Redis/gRPC)
		IsolationOK: true,
		LastChecked: time.Now(),
	}
	e.mu.Lock()
	e.tenants[tenantID] = ctx
	e.mu.Unlock()
	return ctx
}

// EnforceIsolation valida, em memória, se é seguro deixar a requisição
// prosseguir para o tenant informado. Thread-safe via RWMutex.
func (e *IsolationEngine) EnforceIsolation(tenantID string) bool {
	if tenantID == "" {
		return false
	}

	e.mu.RLock()
	ctx, found := e.tenants[tenantID]
	e.mu.RUnlock()

	if !found || time.Since(ctx.LastChecked) > e.cacheTTL {
		ctx = e.register(tenantID)
	}

	return ctx.Active && ctx.IsolationOK
}

// defaultEngine é o singleton consumido pelo guardian_middleware.go.
var defaultEngine = NewIsolationEngine(30 * time.Second)

// EnforceIsolation é o ponto de entrada usado pelo middleware HTTP.
func EnforceIsolation(tenantID string) bool {
	return defaultEngine.EnforceIsolation(tenantID)
}