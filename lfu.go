package cache_elimination

import "container/heap"

const CacheLFU = "lfu"

type LFU struct {
	q *Queue
	m map[string]*KeyValue

	*BaseCache
}

func newLFUCache(cache *BaseCache) *LFU {
	q := make(Queue, 0, cache.Capacity)
	return &LFU{
		q:         &q,
		m:         make(map[string]*KeyValue, cache.Capacity),
		BaseCache: cache,
	}
}

func (c *LFU) Get(key string) interface{} {
	e, ok := c.m[key]
	if ok {
		c.q.update(e, e.Value, e.weight+1)
		return e.Value
	}
	return nil
}

func (c *LFU) Set(key string, value interface{}) {
	// 已缓存
	if e, ok := c.m[key]; ok {
		c.q.update(e, e.Value, e.weight+1)
		return
	}

	// 未缓存则缓存
	kv := &KeyValue{
		Key:   key,
		Value: value,
	}
	heap.Push(c.q, kv)
	c.m[key] = kv

	// 超过容量
	if c.q.Len() > c.Capacity {
		e := heap.Pop(c.q)
		delete(c.m, e.(*KeyValue).Key)
	}
}

func (c *LFU) Del(key string) {
	e, ok := c.m[key]
	if ok {
		heap.Remove(c.q, e.index)
		delete(c.m, e.Key)
	}
}

func (c *LFU) Len() int {
	return c.q.Len()
}
