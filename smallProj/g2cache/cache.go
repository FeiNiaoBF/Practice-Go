package g2cache

import (
	"g2cache/lru"
	"sync"
)

// Cache is a thread-safe lru cache
type cache struct {
	mu         sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

func (c *cache) add(key lru.Key, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = lru.New(int(c.cacheBytes), nil)
	}
	c.lru.Add(key, value)
}

func (c *cache) get(key lru.Key) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}
	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}
	return
}
