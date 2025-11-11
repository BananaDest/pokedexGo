// Package pokecache in charge of cache
package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]CacheEntry
	mu      *sync.RWMutex
}
type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = CacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if cacheEntry, ok := c.entries[key]; ok {
		return cacheEntry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			c.mu.Lock()
			for key, cacheEntry := range c.entries {
				if time.Since(cacheEntry.createdAt) > interval {
					delete(c.entries, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{entries: map[string]CacheEntry{}, mu: &sync.RWMutex{}}
	c.reapLoop(interval)
	return &c
}
