package pokecache

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
	mu           *sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	cacheEntry := CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.mu.Lock()
	c.cacheEntries[key] = cacheEntry
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cacheEntry, found := c.cacheEntries[key]
	return cacheEntry.val, found
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.cacheEntries {
		if entry.createdAt.Before(now.Add(-last)) {
			delete(c.cacheEntries, key)
		}
	}
}

func NewCache(interval time.Duration) Cache {
	mutex := sync.Mutex{}
	cache := Cache{
		cacheEntries: map[string]CacheEntry{},
		mu:           &mutex,
	}
	go cache.reapLoop(interval)
	return cache
}
