package simple_cache

import (
	"sync"
)

type Cache struct {
	sync.RWMutex
	items map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.Lock()
	defer c.Unlock()

	c.items[key] = value
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()

	item, found := c.items[key]

	return item, found
}

func (c *Cache) Delete(key string) {
	c.Lock()
	defer c.Unlock()

	if _, found := c.items[key]; found {
		delete(c.items, key)
	}
}
