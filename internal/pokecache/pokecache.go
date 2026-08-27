package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

type cache struct {
	Entries  map[string]cacheEntry
	mu       *sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *cache {
	c := &cache{
		Entries:  make(map[string]cacheEntry),
		mu:       &sync.Mutex{},
		interval: interval,
	}

	go c.reapLoop()

	return c
}

func (c *cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Entries[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c *cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.Entries[key]
	return entry.val, ok
}

func (c *cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.Entries {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.Entries, key)
			}
		}
		c.mu.Unlock()
	}
}
