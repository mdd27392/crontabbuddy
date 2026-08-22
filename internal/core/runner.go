package core

import (
	"context"
	"time"
)

// Runner executes a series of steps with timing and failure reporting.
type Runner struct {
	logger *Logger
	steps []func(ctx context.Context) error
}

// NewRunner builds a runner bound to a logger.
func NewRunner(logger *Logger) *Runner {
	return &Runner{logger: logger}
}

// Add appends a step to the runner.
func (r *Runner) Add(step func(ctx context.Context) error) *Runner {
	r.steps = append(r.steps, step)
	return r
}

// Run executes every step, returning the first error encountered.
func (r *Runner) Run(ctx context.Context) error {
	for i, step := range r.steps {
 start := time.Now()
 if err := step(ctx); err != nil {
 r.logger.Errorf("step %d failed after %s: %v", i+1, time.Since(start), err)
 return err
 }
 if ctx.Err() != nil {
 return ctx.Err()
 }
 r.logger.Infof("step %d done in %s", i+1, time.Since(start))
	}
	return nil
}