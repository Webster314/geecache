package lru

import (
	"container/list"
)

type Cache struct{
    maxBytes   uint64;
    nBytes     uint64;
    ll         *list.List;
    cache      map[string]*list.Element;
    onEvicted  func(string, Value);
}


type Value interface{
    Len() int
}

type Entry struct{
    k string;
    v Value;
}

func New(maxBytes uint64, onEvicted func(string, Value)) (*Cache){
    return &Cache{
        maxBytes   : maxBytes,
        ll         : list.New(),
        cache      : make(map[string]*list.Element),
        onEvicted  : onEvicted,
    }
}

func (c * Cache)Len() int{
    return c.ll.Len()
}

func (c *Cache)Get(k string) (value Value, ok bool){
    if ele, ok := c.cache[k]; ok{
        kv := ele.Value.(*Entry)
        c.ll.MoveToFront(ele)
        return kv.v, true 
    }
    return
}

func (c *Cache)RemoveOldest(){
    if ele := c.ll.Back(); ele != nil{
        c.ll.Remove(ele)
        kv := ele.Value.(*Entry)
        delete(c.cache, kv.k)
        c.nBytes -= uint64(len(kv.k)) + uint64(kv.v.Len())
        if c.onEvicted != nil{
            c.onEvicted(kv.k, kv.v)
        }
    }
}

func (c *Cache)Add(k string, v Value){
    if ele, ok := c.cache[k]; ok{
        c.ll.MoveToFront(ele)
        kv := ele.Value.(*Entry)
        kv.v = v
        c.nBytes += (uint64(v.Len()) - uint64(kv.v.Len()))
    }else{
        kv := Entry{ k : k, v : v }
        ele := c.ll.PushFront(&kv)
        c.cache[k] = ele
        c.nBytes += uint64(len(k)) + uint64(v.Len())
    }
    for c.maxBytes != 0 && c.maxBytes < c.nBytes{
        c.RemoveOldest()
    }
}
