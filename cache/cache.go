package cache

import (
	"sync"
	"time"
)

type Cache struct {
	data      interface{}
	timestamp time.Time
	ttl       time.Duration
	mutex     sync.RWMutex
}

func NewCache(ttl time.Duration) *Cache {
	return &Cache{ttl: ttl}
}

func (c *Cache) Get() (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if time.Since(c.timestamp) > c.ttl {
		return nil, false
	}
	return c.data, true
}

func (c *Cache) Set(data interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data = data
	c.timestamp = time.Now()
}
