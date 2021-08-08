package cache_elimination

type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	Del(key string)
	Len() int
}

type BaseCache struct {
	CacheType string
	Capacity  int
}

func (c *BaseCache) Set(key string, value interface{}) {}

func (c *BaseCache) Get(key string) interface{} {
	return nil
}

func (c *BaseCache) Del(key string) {}

func (c *BaseCache) Len() int {
	return 0
}

func NewCache(cacheType string, capacity int) Cache {
	baseCache := &BaseCache{
		CacheType: cacheType,
		Capacity:  capacity,
	}
	switch cacheType {
	case CacheSimple:
		return newSimpleCache(baseCache)
	case CacheFIFO:
		return newFifoCache(baseCache)
	case CacheLRU:
		return newLRUCache(baseCache)
	}
	return newSimpleCache(baseCache)
}
