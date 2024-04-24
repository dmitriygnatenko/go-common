package memory_cache

import (
	"sync"
)

type Cache struct {
	mu    sync.RWMutex
	items map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = value
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]

	return item, found
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, found := c.items[key]; found {
		delete(c.items, key)
	}
}
