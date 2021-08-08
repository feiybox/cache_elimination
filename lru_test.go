package cache_elimination

import "testing"

func TestLRU_Get(t *testing.T) {
	c := NewCache(CacheLRU, 2)
	t.Log(c.Get("key"))
}

func TestLRU_Set(t *testing.T) {
	c := NewCache(CacheLRU, 2)
	c.Set("key", 1)
	if c.Get("key") != 1 {
		t.Error(c.Get("key"))
	}
	c.Set("key1", 1)
	c.Set("key2", 1)
	if c.Get("key") != nil {
		t.Error(c.Get("key"))
	}
	if c.Get("key1") != 1 {
		t.Error(c.Get("key1"))
	}
}

func TestLRU_Del(t *testing.T) {
	c := NewCache(CacheLRU, 2)
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

func TestLRU_Len(t *testing.T) {
	c := NewCache(CacheLRU, 2)
	if c.Len() != 0 {
		t.Error("len:", c.Len())
	}
	c.Set("key", 1)
	if c.Len() != 1 {
		t.Error("len:", c.Len())
	}
	c.Set("key1", 1)
	c.Set("key2", 1)
	if c.Len() != 2 {
		t.Error("len:", c.Len())
	}
}
