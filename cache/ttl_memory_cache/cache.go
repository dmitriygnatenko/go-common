package ttl_memory_cache

import (
	"sync"
	"time"
)

type Cache struct {
	expiration *time.Duration
	mu         sync.RWMutex
	items      map[string]Item
}

type Item struct {
	Value     interface{}
	ExpiredAt *time.Time
}

func NewCache(c Config) *Cache {
	cache := Cache{
		items:      make(map[string]Item),
		expiration: c.expiration,
	}

	return &cache
}

func (c *Cache) Set(key string, value interface{}, expiration *time.Duration) {
	item := Item{
		Value: value,
	}

	if expiration != nil {
		expiredAt := time.Now().Add(*expiration)
		item.ExpiredAt = &expiredAt
	} else if c.expiration != nil {
		expiredAt := time.Now().Add(*c.expiration)
		item.ExpiredAt = &expiredAt
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = item
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()

	item, found := c.items[key]
	if !found {
		c.mu.RUnlock()
		return nil, false
	}

	if item.ExpiredAt == nil {
		c.mu.RUnlock()
		return item.Value, true
	}

	if time.Now().Before(*item.ExpiredAt) {
		c.mu.RUnlock()
		return item.Value, true
	}

	c.mu.RUnlock() // defer not used due to deadlocks

	c.Delete(key)

	return nil, false
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, found := c.items[key]; found {
		delete(c.items, key)
	}
}

func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[string]Item)
}
