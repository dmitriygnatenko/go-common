package lru_memory_cache

import (
	"container/list"
	"sync"
)

type Cache[K comparable, V any] struct {
	capacity int

	mu    sync.RWMutex
	items map[K]*list.Element
	queue *list.List
}

type Item[K comparable, V any] struct {
	Key   K
	Value V
}

func NewCache[K comparable, V any](capacity int) *Cache[K, V] {
	return &Cache[K, V]{
		capacity: capacity,
		items:    make(map[K]*list.Element),
		queue:    list.New(),
	}
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if element, found := c.items[key]; found {
		c.queue.MoveToFront(element)
		element.Value.(*Item[K, V]).Value = value
		return
	}

	if c.queue.Len() == c.capacity {
		c.clean()
	}

	item := &Item[K, V]{
		Key:   key,
		Value: value,
	}

	element := c.queue.PushFront(item)
	c.items[item.Key] = element
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if element, found := c.items[key]; found {
		c.queue.MoveToFront(element)
		return element.Value.(*Item[K, V]).Value, true
	}

	return Item[K, V]{}.Value, false
}

func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if element, found := c.items[key]; found {
		c.queue.Remove(element)
		delete(c.items, key)
	}
}

func (c *Cache[K, V]) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[K]*list.Element)
	c.queue = list.New()
}

func (c *Cache[K, V]) clean() {
	if element := c.queue.Back(); element != nil {
		item := c.queue.Remove(element).(*Item[K, V])
		delete(c.items, item.Key)
	}
}
