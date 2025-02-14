package ttl_memory_cache

import (
	"sync"
	"time"
)

type Cache[K comparable, V any] struct {
	expiration time.Duration
	mu         sync.RWMutex
	items      map[K]Item[V]
}

type Item[V any] struct {
	Value     V
	ExpiredAt time.Time
}

func NewCache[K comparable, V any](defaultExpiration time.Duration) *Cache[K, V] {
	cache := Cache[K, V]{
		items:      make(map[K]Item[V]),
		expiration: defaultExpiration,
	}

	return &cache
}

func (c *Cache[K, V]) Set(key K, value V, expiration *time.Duration) {
	var expiredAt time.Time

	if expiration != nil {
		expiredAt = time.Now().Add(*expiration)
	} else {
		expiredAt = time.Now().Add(c.expiration)
	}

	item := Item[V]{
		Value:     value,
		ExpiredAt: expiredAt,
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = item
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()

	item, found := c.items[key]
	if !found {
		c.mu.RUnlock()
		return item.Value, false
	}

	if time.Now().Before(item.ExpiredAt) {
		c.mu.RUnlock()
		return item.Value, true
	}

	c.mu.RUnlock() // defer not used due to deadlocks

	c.Delete(key)

	return Item[V]{}.Value, false
}

func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, found := c.items[key]; found {
		delete(c.items, key)
	}
}

func (c *Cache[K, V]) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[K]Item[V])
}
