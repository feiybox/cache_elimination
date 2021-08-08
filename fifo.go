package cache_elimination

import "container/list"

const CacheFIFO = "fifo"

// 先进先出算法
type Fifo struct {
	ll *list.List
	m  map[string]*list.Element

	*BaseCache
}

func newFifoCache(cache *BaseCache) *Fifo {
	return &Fifo{
		ll:        list.New(),
		m:         make(map[string]*list.Element, cache.Capacity),
		BaseCache: cache,
	}
}

func (c *Fifo) Set(key string, value interface{}) {
	// 已缓存
	if e, ok := c.m[key]; ok {
		c.ll.MoveToBack(e)
		return
	}

	// 未缓存则缓存
	kv := &KeyValue{
		Key:   key,
		Value: value,
	}
	e := c.ll.PushBack(kv)
	c.m[key] = e

	// 超过容量
	if c.ll.Len() > c.Capacity {
		e := c.ll.Front()
		delete(c.m, e.Value.(*KeyValue).Key)
		c.ll.Remove(e)
	}
}

func (c *Fifo) Get(key string) interface{} {
	e, ok := c.m[key]
	if ok {
		return e.Value.(*KeyValue).Value
	}
	return nil
}

func (c *Fifo) Del(key string) {
	e, ok := c.m[key]
	if ok {
		c.ll.Remove(e)
		delete(c.m, key)
	}
}

func (c *Fifo) Len() int {
	return c.ll.Len()
}
