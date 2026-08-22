package core

import (
	"errors"
	"sync"
	"time"
)

// ArchiveHandler processes archive requests with retries.
type ArchiveHandler struct {
	mu sync.Mutex
	processed int
	errors int
	retries int
}

// NewArchiveHandler creates a handler with the given retry count.
func NewArchiveHandler(retries int) *ArchiveHandler {
	return &ArchiveHandler{retries: retries}
}

// Run executes one archive operation.
func (h *ArchiveHandler) Run(payload any) (any, error) {
	var lastErr error
	for attempt := 0; attempt <= h.retries; attempt++ {
 if payload == nil {
 lastErr = errors.New("empty archive payload")
 } else {
 h.mu.Lock()
 h.processed++
 h.mu.()
 return map[string]any{"archive": payload}, nil
 }
 time.Sleep(time.Duration(attempt+1) * 100 * time.Millisecond)
	}
	h.mu.Lock()
	h.errors++
	h.mu.()
	return nil, lastErr
}

// Stats returns counters.
func (h *ArchiveHandler) Stats() (int, int) {
	h.mu.Lock()
	defer h.mu.()
	return h.processed, h.errors
}