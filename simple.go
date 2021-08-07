package cache_elimination

const CacheSimple = "simple"

type Simple struct {
	m map[string]interface{}

	*BaseCache
}

func newSimpleCache(cache *BaseCache) *Simple {
	return &Simple{
		m:         make(map[string]interface{}, cache.Capacity),
		BaseCache: cache,
	}
}

func (c *Simple) Set(key string, value interface{}) {
	c.m[key] = value
}

func (c *Simple) Get(key string) interface{} {
	return c.m[key]
}

func (c *Simple) Del(key string) {
	delete(c.m, key)
}

func (c *Simple) Len() int {
	return len(c.m)
}
