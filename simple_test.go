package cache_elimination

import "testing"

func TestSimple_Get(t *testing.T) {
	c := NewCache(CacheSimple, 10)
	t.Log(c.Get("key"))
}

func TestSimple_Set(t *testing.T) {
	c := NewCache(CacheSimple, 10)
	c.Set("key", 1)
	if c.Get("key") != 1 {
		t.Error(c.Get("key"))
	}
}

func TestSimple_Del(t *testing.T) {
	c := NewCache(CacheSimple, 10)
	c.Del("key")
	c.Set("key", 1)
	if c.Get("key") != 1 {
		t.Error(c.Get("key"))
	}
	c.Del("key")
	if c.Get("key") != nil {
		t.Error(c.Get("key"))
	}
}

func TestSimple_Len(t *testing.T) {
	c := NewCache(CacheSimple, 10)
	if c.Len() != 0 {
		t.Error("len:", c.Len())
	}
	c.Set("key", 1)
	if c.Len() != 1 {
		t.Error("len:", c.Len())
	}
}
