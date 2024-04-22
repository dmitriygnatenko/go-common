package ttl_memory_cache

import (
	"sync"
	"time"
)

type Cache struct {
	expiration      *time.Duration
	cleanupInterval *time.Duration

	sync.RWMutex
	items map[string]Item
}

type Item struct {
	Value     interface{}
	ExpiredAt *time.Time
}

func NewCache(c Config) *Cache {
	cache := Cache{
		items:           make(map[string]Item),
		expiration:      c.expiration,
		cleanupInterval: c.cleanupInterval,
	}

	if c.cleanupInterval != nil {
		go cache.cleanupWorker()
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

	c.Lock()
	defer c.Unlock()

	c.items[key] = item
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, false
	}

	if item.ExpiredAt != nil && time.Now().After(*item.ExpiredAt) {
		c.Delete(key)
		return nil, false
	}

	return item.Value, true
}

func (c *Cache) Delete(key string) {
	c.Lock()
	defer c.Unlock()

	if _, found := c.items[key]; found {
		delete(c.items, key)
	}
}

func (c *Cache) cleanupWorker() {
	for {
		<-time.After(*c.cleanupInterval)

		if len(c.items) == 0 {
			return
		}

		now := time.Now()
		for key, item := range c.items {
			if item.ExpiredAt != nil && now.After(*item.ExpiredAt) {
				c.Delete(key)
			}
		}
	}
}
