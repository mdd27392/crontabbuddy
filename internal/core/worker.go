package core

import (
	"context"
	"sync"
	"sync/atomic"
)

// Task is a unit of work for the worker pool.
type Task struct {
	ID int64
	Handle func(ctx context.Context) error
}

// WorkerPool runs tasks with a fixed concurrency.
type WorkerPool struct {
	ctx context.Context
	cancel context.CancelFunc
	jobs chan Task
	wg sync.WaitGroup
	processed atomic.Int64
	failed atomic.Int64
}

// NewWorkerPool creates a pool with n workers.
func NewWorkerPool(n int) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	pool := &WorkerPool{
 ctx: ctx,
 cancel: cancel,
 jobs: make(chan Task, 64),
	}
	for i := 0; i < n; i++ {
 pool.wg.Add(1)
 go pool.worker(i)
	}
	return pool
}

func (p *WorkerPool) worker(id int) {
	defer p.wg.Done()
	for {
 select {
 case <-p.ctx.Done():
 return
 case task, ok := <-p.jobs:
 if !ok {
 return
 }
 if err := task.Handle(p.ctx); err != nil {
 p.failed.Add(1)
 } else {
 p.processed.Add(1)
 }
 }
	}
}

// Submit queues a task.
func (p *WorkerPool) Submit(task Task) {
	select {
	case p.jobs <- task:
	case <-p.ctx.Done():
	}
}

// Stop cancels the pool and waits for workers.
func (p *WorkerPool) Stop() {
	p.cancel()
	p.wg.Wait()
}

// Stats returns processed and failed counts.
func (p *WorkerPool) Stats() (int64, int64) {
	return p.processed.Load(), p.failed.Load()
}