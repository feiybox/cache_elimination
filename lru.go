package cache_elimination

import "container/list"

const CacheLRU = "lru"

type LRU struct {
	ll *list.List // 链表最后的节点为最新使用的节点
	m  map[string]*list.Element

	*BaseCache
}

func newLRUCache(cache *BaseCache) *LRU {
	return &LRU{
		ll:        list.New(),
		m:         make(map[string]*list.Element, cache.Capacity),
		BaseCache: cache,
	}
}

func (c *LRU) Get(key string) interface{} {
	if e, ok := c.m[key]; ok {
		c.ll.MoveToBack(e)
		return e.Value.(*KeyValue).Value
	}
	return nil
}

func (c *LRU) Set(key string, value interface{}) {
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

func (c *LRU) Del(key string) {
	e, ok := c.m[key]
	if ok {
		c.ll.Remove(e)
		delete(c.m, e.Value.(*KeyValue).Key)
	}
}

func (c *LRU) Len() int {
	return c.ll.Len()
}
