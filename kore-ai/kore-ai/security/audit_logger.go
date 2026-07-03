package security

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"
)

// ViolationEvent representa um registro de tentativa de acesso bloqueada.
type ViolationEvent struct {
	Timestamp time.Time `json:"timestamp"`
	IP        string    `json:"ip"`
	Path      string    `json:"path"`
	Reason    string    `json:"reason"`
}

const logChannelBufferSize = 1024

// AuditLogger grava eventos de violação de forma assíncrona via canal +
// goroutine dedicada, evitando que o hot path do proxy fique bloqueado
// por I/O de log.
type AuditLogger struct {
	eventChan chan ViolationEvent
	logger    *log.Logger
	wg        sync.WaitGroup
}

// NewAuditLogger cria um logger e inicia o worker consumidor.
func NewAuditLogger(output *os.File) *AuditLogger {
	if output == nil {
		output = os.Stdout
	}
	al := &AuditLogger{
		eventChan: make(chan ViolationEvent, logChannelBufferSize),
		logger:    log.New(output, "", 0),
	}
	al.wg.Add(1)
	go al.worker()
	return al
}

func (al *AuditLogger) worker() {
	defer al.wg.Done()
	for event := range al.eventChan {
		data, err := json.Marshal(event)
		if err != nil {
			al.logger.Printf(`{"level":"error","msg":"falha ao serializar evento de auditoria: %v"}`, err)
			continue
		}
		al.logger.Println(string(data))
	}
}

// LogViolation envia o evento para o canal sem bloquear o caller.
// Em caso de canal saturado (burst de ataques), o evento é descartado
// e um aviso síncrono é emitido — prioridade é nunca degradar a latência
// do proxy por causa do logging.
func (al *AuditLogger) LogViolation(ip, path, reason string) {
	event := ViolationEvent{
		Timestamp: time.Now().UTC(),
		IP:        ip,
		Path:      path,
		Reason:    reason,
	}
	select {
	case al.eventChan <- event:
	default:
		al.logger.Printf(`{"level":"warn","msg":"canal de auditoria saturado, evento descartado","ip":"%s","path":"%s"}`, ip, path)
	}
}

// Close drena o canal e aguarda o worker finalizar. Deve ser chamado no
// shutdown gracioso do servidor (defer security.Close() no main).
func (al *AuditLogger) Close() {
	close(al.eventChan)
	al.wg.Wait()
}

// defaultLogger é o singleton usado pela função de pacote LogViolation,
// consumido diretamente pelo guardian_middleware.go.
var defaultLogger = NewAuditLogger(os.Stdout)

// LogViolation é o ponto de entrada usado pelo middleware HTTP.
func LogViolation(ip, path, reason string) {
	defaultLogger.LogViolation(ip, path, reason)
}

// Close encerra o logger padrão (chamar no shutdown da aplicação).
func Close() {
	defaultLogger.Close()
}