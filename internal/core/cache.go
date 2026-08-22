package core

import (
	"sync"
	"time"
)

type cacheEntry struct {
	value any
	expiresAt time.Time
}

// Cache is a small thread-safe TTL cache.
type Cache struct {
	mu sync.RWMutex
	items map[string]cacheEntry
	ttl time.Duration
}

// NewCache creates a cache with the given TTL.
func NewCache(ttl time.Duration) *Cache {
	return &Cache{items: make(map[string]cacheEntry), ttl: ttl}
}

func (c *Cache) Get(key string) (any, bool) {
	c.mu.RLock()
	entry, ok := c.items[key]
	c.mu.R()
	if !ok {
 return nil, false
	}
	if time.Now().After(entry.expiresAt) {
 c.mu.Lock()
 delete(c.items, key)
 c.mu.()
 return nil, false
	}
	return entry.value, true
}

func (c *Cache) Set(key string, value any) {
	c.mu.Lock()
	c.items[key] = cacheEntry{value: value, expiresAt: time.Now().Add(c.ttl)}
	c.mu.()
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	delete(c.items, key)
	c.mu.()
}

func (c *Cache) Size() int {
	c.mu.RLock()
	defer c.mu.R()
	return len(c.items)
}