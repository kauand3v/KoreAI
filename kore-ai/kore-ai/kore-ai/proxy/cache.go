// kore-ai/proxy/cache.go
package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

// SemanticCache define o contrato para interceptação de prompts repetidos/similares.
// Implementações alternativas (ex: cache vetorial) podem ser plugadas sem alterar o resto do gateway.
type SemanticCache interface {
	CheckCache(tenantID, prompt string) (string, bool)
	SaveCache(tenantID, prompt, response string) error
}

// RedisSemanticCache implementa SemanticCache usando Redis como backend distribuído.
// A "semântica" aqui é obtida via normalização do prompt (lowercase, trim, colapso de
// espaços) antes do hashing -- garante hits para variações triviais de escrita.
// Para similaridade vetorial real, plugar um pipeline de embeddings antes do hash.
type RedisSemanticCache struct {
	client *redis.Client
	ttl    time.Duration
}

// NewRedisSemanticCache cria o client Redis com timeouts agressivos -- o cache nunca
// pode se tornar o gargalo de latência do gateway.
func NewRedisSemanticCache(addr string, ttl time.Duration) *RedisSemanticCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:         addr,
		DialTimeout:  2 * time.Second,
		ReadTimeout:  500 * time.Millisecond,
		WriteTimeout: 500 * time.Millisecond,
		PoolSize:     50,
	})
	return &RedisSemanticCache{client: rdb, ttl: ttl}
}

var whitespaceRegex = regexp.MustCompile(`\s+`)

// normalize prepara o prompt para o hashing semântico simplificado.
func normalize(prompt string) string {
	p := strings.ToLower(strings.TrimSpace(prompt))
	return whitespaceRegex.ReplaceAllString(p, " ")
}

// cacheKey gera uma chave determinística isolada por tenant (multi-tenancy seguro).
func cacheKey(tenantID, prompt string) string {
	hash := sha256.Sum256([]byte(normalize(prompt)))
	return fmt.Sprintf("kore:cache:%s:%s", tenantID, hex.EncodeToString(hash[:]))
}

// CheckCache consulta o Redis com timeout curto -- nunca deve travar o hot-path do gateway.
func (r *RedisSemanticCache) CheckCache(tenantID, prompt string) (string, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	val, err := r.client.Get(ctx, cacheKey(tenantID, prompt)).Result()
	if err != nil {
		if err != redis.Nil {
			log.Printf("[cache] erro ao consultar redis: %v", err)
		}
		return "", false
	}
	return val, true
}

// SaveCache grava a resposta no Redis com TTL configurado. O chamador deve disparar
// esta função em uma goroutine separada para não bloquear a resposta ao cliente.
func (r *RedisSemanticCache) SaveCache(tenantID, prompt, response string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	if err := r.client.Set(ctx, cacheKey(tenantID, prompt), response, r.ttl).Err(); err != nil {
		return fmt.Errorf("falha ao salvar cache: %w", err)
	}
	return nil
}