package lru_memory_cache

import (
	"container/list"
	"errors"
	"sync"
)

type Cache struct {
	capacity int

	mu    sync.RWMutex
	items map[string]*list.Element
	queue *list.List
}

type Item struct {
	Key   string
	Value interface{}
}

func NewCache(c Config) (*Cache, error) {
	if c.capacity == 0 {
		return nil, errors.New("empty capacity")
	}

	return &Cache{
		capacity: int(c.capacity),
		items:    make(map[string]*list.Element),
		queue:    list.New(),
	}, nil
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if element, found := c.items[key]; found {
		c.queue.MoveToFront(element)
		element.Value.(*Item).Value = value
		return
	}

	if c.queue.Len() == c.capacity {
		c.clean()
	}

	item := &Item{
		Key:   key,
		Value: value,
	}

	element := c.queue.PushFront(item)
	c.items[item.Key] = element
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if element, found := c.items[key]; found {
		c.queue.MoveToFront(element)
		return element.Value.(*Item).Value, true
	}

	return nil, false
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if element, found := c.items[key]; found {
		c.queue.Remove(element)
		delete(c.items, key)
	}
}

func (c *Cache) clean() {
	if element := c.queue.Back(); element != nil {
		item := c.queue.Remove(element).(*Item)
		delete(c.items, item.Key)
	}
}
