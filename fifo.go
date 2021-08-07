package cache_elimination

const CacheFIFO = "fifo"

// 先进先出算法
type Fifo struct {
	*BaseCache
}

func newFifoCache(cache *BaseCache) *Fifo {
	return &Fifo{
		BaseCache: cache,
	}
}

func (c *Fifo) Set(key string, value interface{}) {
	// todo
}

func (c *Fifo) Get(key string) interface{} {
	// todo
	return nil
}

func (c *Fifo) Del(key string) {
	// todo
}

func (c *Fifo) Len() int {
	// todo
	return 0
}
