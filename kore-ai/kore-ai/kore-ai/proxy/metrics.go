// kore-ai/proxy/metrics.go
package main

import (
	"log"
	"sync"
	"time"

	"kore-ai/core/types"
)

// MetricsCollector recebe payloads via channel e processa de forma assíncrona,
// garantindo que o hot-path de requisições nunca seja bloqueado por I/O de observabilidade.
type MetricsCollector struct {
	ch chan types.MetricsPayload
	wg sync.WaitGroup
}

// NewMetricsCollector inicia N workers consumindo o channel em paralelo.
func NewMetricsCollector(bufferSize, workers int) *MetricsCollector {
	mc := &MetricsCollector{
		ch: make(chan types.MetricsPayload, bufferSize),
	}

	for i := 0; i < workers; i++ {
		mc.wg.Add(1)
		go mc.worker(i)
	}

	return mc
}

// worker drena o channel até ele ser fechado (graceful shutdown via Shutdown()).
func (mc *MetricsCollector) worker(id int) {
	defer mc.wg.Done()
	for payload := range mc.ch {
		mc.process(payload)
	}
}

// process simula o envio para um backend de observabilidade (ex: Prometheus/Grafana Loki).
func (mc *MetricsCollector) process(p types.MetricsPayload) {
	// TODO: substituir por exportador real (OTLP/Prometheus pushgateway).
	log.Printf(
		"[metrics] tenant=%s model=%s latency_ms=%d tokens=%d cache_hit=%t err=%q ts=%s",
		p.TenantID, p.Model, p.LatencyMs, p.TokensUsed, p.CacheHit, p.Error, p.Timestamp.Format(time.RFC3339),
	)
}

// Publish envia a métrica sem bloquear o caller. Se o buffer estiver cheio, descarta
// o evento (drop-on-full) -- proteger a latência da requisição é sempre prioridade.
func (mc *MetricsCollector) Publish(p types.MetricsPayload) {
	select {
	case mc.ch <- p:
	default:
		log.Printf("[metrics] buffer cheio, descartando métrica do tenant=%s", p.TenantID)
	}
}

// Shutdown realiza o graceful shutdown, aguardando os workers drenarem o buffer pendente.
func (mc *MetricsCollector) Shutdown() {
	close(mc.ch)
	mc.wg.Wait()
}