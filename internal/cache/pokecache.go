package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]cacheEntry
	mu   *sync.Mutex
}
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		data: make(map[string]cacheEntry),
		mu:   &sync.Mutex{},
	}
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			c.reapLoop(interval)
		}
	}()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Lock()
	c.data[key] = entry
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v, exists := c.data[key]; exists {
		return v.val, true
	} else {
		return nil, false
	}

}

func (c *Cache) reapLoop(interval time.Duration) {
	c.mu.Lock()
	for k, v := range c.data {
		if time.Since(v.createdAt) > interval {
			delete(c.data, k)
		}
	}
	c.mu.Unlock()
}
