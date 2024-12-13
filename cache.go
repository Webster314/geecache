package geecache

import (
	"geecache/lru"
	"sync"
)

type cache struct {
	mtx        sync.Mutex
	lru        *lru.Cache
	cacheBytes uint64
}

func (c *cache) add(k string, v ByteView) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}
	c.lru.Add(k, v)
}

func (c *cache) get(k string) (ByteView, bool) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	if c.lru == nil {
        return ByteView{}, false
	}
    if v, ok := c.lru.Get(k); ok {
		return v.(ByteView), ok
	}
	return ByteView{}, false
}
