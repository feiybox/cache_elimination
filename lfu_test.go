package cache_elimination

import "testing"

func TestLFU_Get(t *testing.T) {
	c := NewCache(CacheLFU, 2)
	t.Log(c.Get("key"))
}

func TestLFU_Set(t *testing.T) {
	c := NewCache(CacheLFU, 2)
	c.Set("key", 1)
	c.Set("key1", 1)
	c.Get("key1")
	c.Set("key2", 1)
	if c.Get("key") != nil {
		t.Error(c.Get("key"))
	}
	if c.Get("key1") != 1 {
		t.Error(c.Get("key1"))
	}
}

func TestLFU_Del(t *testing.T) {
	c := NewCache(CacheLFU, 2)
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

func TestLFU_Len(t *testing.T) {
	c := NewCache(CacheLFU, 2)
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
