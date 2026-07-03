package security

import (
	"context"
	"encoding/json"
	"net/http"
)

type contextKey string

// TenantIDContextKey é a chave usada para injetar o tenantID validado
// no context.Context, consumível por qualquer handler downstream.
const TenantIDContextKey contextKey = "tenantID"

const tenantHeader = "x-tenant-id"

type errorResponse struct {
	Error   string `json:"error"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func writeForbidden(w http.ResponseWriter, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	_ = json.NewEncoder(w).Encode(errorResponse{
		Error:   "forbidden",
		Code:    code,
		Message: message,
	})
}

func clientIP(r *http.Request) string {
	if fwd := r.Header.Get("X-Forwarded-For"); fwd != "" {
		return fwd
	}
	return r.RemoteAddr
}

// GuardianMiddleware é a primeira linha de defesa do KoreAI contra
// cross-tenant data leaks: roda antes de qualquer handler de negócio.
//
//  1. Sem header x-tenant-id  -> audita e responde 403 imediatamente.
//  2. Header presente         -> valida via isolation_engine.
//  3. Aprovado                -> injeta tenantID no contexto e segue.
func GuardianMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tenantID := r.Header.Get(tenantHeader)

		if tenantID == "" {
			LogViolation(clientIP(r), r.URL.Path, "header x-tenant-id ausente")
			writeForbidden(w, "TENANT_MISSING", "Requisição rejeitada: o header x-tenant-id é obrigatório.")
			return
		}

		if !EnforceIsolation(tenantID) {
			LogViolation(clientIP(r), r.URL.Path, "falha no isolation_engine para tenant "+tenantID)
			writeForbidden(w, "ISOLATION_VIOLATION", "Requisição rejeitada: falha na validação de isolamento do tenant.")
			return
		}

		ctx := context.WithValue(r.Context(), TenantIDContextKey, tenantID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// TenantIDFromContext recupera de forma segura o tenantID injetado pelo
// GuardianMiddleware. Use isto nos handlers downstream em vez de ler o
// header novamente.
func TenantIDFromContext(ctx context.Context) (string, bool) {
	tenantID, ok := ctx.Value(TenantIDContextKey).(string)
	return tenantID, ok
}