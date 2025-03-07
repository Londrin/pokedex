package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	CreatedAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
	done  chan bool
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
		done:  make(chan bool),
	}
	go newCache.reapLoop(interval)
	return newCache
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		CreatedAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, found := c.cache[key]

	if !found {
		return []byte{}, false
	}

	c.cache[key] = cacheEntry{
		CreatedAt: time.Now(),
		val:       entry.val,
	}

	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			c.mu.Lock()
			for key, entry := range c.cache {
				if time.Since(entry.CreatedAt) > interval {
					delete(c.cache, key)
				}
			}
			c.mu.Unlock()

		case <-c.done:
			return
		}
	}
}

func (c *Cache) Stop() {
	close(c.done)
}
