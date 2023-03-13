package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (c *lruCache) Set(key Key, value interface{}) (ok bool) {
	val, ok := c.items[key]
	if ok {
		val.Value = value
		c.queue.MoveToFront(val)
		return
	}
	c.items[key] = c.queue.PushFront(value)
	if c.queue.Len() > c.capacity {
		for key, item := range c.items {
			if item == c.queue.Back() {
				delete(c.items, key)
			}
		}
		c.queue.Remove(c.queue.Back())
	}
	return
}

func (c *lruCache) Get(key Key) (value interface{}, ok bool) {
	val, ok := c.items[key]
	if ok {
		value = val.Value
		c.queue.MoveToFront(val)
	}
	return
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
