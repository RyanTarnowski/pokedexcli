package pokeapi

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheEntries map[string]CacheEntry
	mux          sync.Mutex
	interval     time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		cacheEntries: map[string]CacheEntry{},
		mux:          sync.Mutex{},
		interval:     interval,
	}
	go cache.reapLoop()
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.cacheEntries[key] = CacheEntry{createdAt: time.Now().UTC(), val: val}
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	entry := c.cacheEntries[key]
	c.mux.Unlock()

	if entry.val == nil {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	for tick := range ticker.C {
		c.mux.Lock()
		for key, ce := range c.cacheEntries {
			if ce.createdAt.Before(tick.UTC().Add(-c.interval)) {
				delete(c.cacheEntries, key)
			}
		}
		c.mux.Unlock()
	}
}
