package geecache

import (
	"log"
	"sync"
)

type Getter interface{
    Get(string) ([]byte, error)
}

type GetterFunc func(string) ([]byte, error)

func (gf GetterFunc)Get(k string) ([]byte, error){
    return gf(k)
}

type Group struct{
    name       string
    mainCache  *cache
    getter     Getter
}

var(
    mtx    sync.RWMutex
    groups = make(map[string]*Group)
)

func NewGroup(name string, cacheBytes uint64, getter Getter) (*Group){
    if(getter == nil){
        panic("nil getter")
    }
    mtx.Lock()
    defer mtx.Unlock()
    g := &Group{
        name       : name,
        mainCache  : &cache{cacheBytes : cacheBytes},
        getter     : getter,
    }
    groups[name] = g
    return g
}

func GetGroup(name string) (*Group){
    mtx.RLock()
    g := groups[name]
    mtx.RUnlock()
    return g
}

func (g *Group)Get(k string) (ByteView, error){
    if v, ok := g.mainCache.get(k); ok {
        log.Println("[GEECACHE] hit")
        return v, nil
    }
    return g.load(k)
}

func (g * Group)load(k string) (ByteView, error){
    return g.loadLocally(k)
}

func (g * Group)loadLocally(k string) (ByteView, error){
    b, e := g.getter.Get(k)
    if e != nil{
        return ByteView{}, e
    }
    v := ByteView{b : CloneBytes(b)}
    g.populateCache(k, v)
    return v, nil
}

func (g * Group)populateCache(k string, v ByteView){
    g.mainCache.add(k, v)
}


