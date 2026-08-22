package core

import (
	"errors"
	"sync"
	"time"
)

// RenderHandler processes render requests with retries.
type RenderHandler struct {
	mu sync.Mutex
	processed int
	errors int
	retries int
}

// NewRenderHandler creates a handler with the given retry count.
func NewRenderHandler(retries int) *RenderHandler {
	return &RenderHandler{retries: retries}
}

// Run executes one render operation.
func (h *RenderHandler) Run(payload any) (any, error) {
	var lastErr error
	for attempt := 0; attempt <= h.retries; attempt++ {
 if payload == nil {
 lastErr = errors.New("empty render payload")
 } else {
 h.mu.Lock()
 h.processed++
 h.mu.()
 return map[string]any{"render": payload}, nil
 }
 time.Sleep(time.Duration(attempt+1) * 100 * time.Millisecond)
	}
	h.mu.Lock()
	h.errors++
	h.mu.()
	return nil, lastErr
}

// Stats returns counters.
func (h *RenderHandler) Stats() (int, int) {
	h.mu.Lock()
	defer h.mu.()
	return h.processed, h.errors
}