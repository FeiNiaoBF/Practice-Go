package lru

import "container/list"

type Cache struct {
	MaxEntries int
	NEntries   int // 当前缓存的项数
	OnEvicted  func(key Key, value any)
	ll         *list.List // 链表的 back 为队头是最久访问的项，front 为队尾, 尾部是最新未被访问的项。
	cache      map[any]*list.Element
}

type Key any

type entry struct {
	key   Key
	value any
}

// New creates a new Cache.
// If maxEntries is zero, the cache has no limit and it is effectively unbounded.
// If onEvicted is nil, no event will be triggered when an entry is deleted.
func New(maxEntries int, onEvicted func(key Key, value any)) *Cache {
	return &Cache{
		MaxEntries: maxEntries,
		NEntries: 0,
		ll:         list.New(),
		cache:      make(map[any]*list.Element),
		OnEvicted:  onEvicted,
	}
}

// Add add and update the cache for a key-value pair.
func (c *Cache) Add(key Key, value any) {
	if c.cache == nil {
		c.cache = make(map[any]*list.Element)
		c.ll = list.New()
	}
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		kv.value = value
		return
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.NEntries++
	}

	if c.MaxEntries != 0 && c.NEntries > c.MaxEntries {
		c.RemoveOldest()
	}
}

func (c *Cache) Get(key Key) (value any, ok bool) {
	if c.cache == nil {
		return
	}
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// Remove removes the provided key from the cache.
func (c *Cache) Remove(key Key) {
	if c.cache == nil {
		return
	}
	if ele, hit := c.cache[key]; hit {
		c.removeElement(ele)
	}
}

// RemoveOldest removes the oldest item from the cache.
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.removeElement(ele)
	}
}

func (c *Cache) removeElement(e *list.Element) {
	c.ll.Remove(e)
	kv := e.Value.(*entry)
	delete(c.cache, kv.key)
	c.NEntries--
	if c.OnEvicted != nil {
		c.OnEvicted(kv.key, kv.value)
	}
}

func (c *Cache) Len() int {
	if c.cache == nil {
		return 0
	}
	return c.NEntries
}

func (c *Cache) Clear() {
	if c.OnEvicted != nil {
		for _, e := range c.cache {
			kv := e.Value.(*entry)
			c.OnEvicted(kv.key, kv.value)
		}
	}
	c.ll = nil
	c.cache = nil
}
