package lru

import (
	"testing"
)

type String string

func (d String)Len() int{
    return len(d)
}

func TestGet(t * testing.T){
    lru := New(0, nil) 
    lru.Add("k1", String("v1"))
    if v, ok := lru.Get("k1"); !ok || v != String("v1") {
        t.Error("cache hit k1 failed")
    }
    if _, ok := lru.Get("k2"); ok{
        t.Error("cache missed k2 failed")
    }
}

func TestAdd(t * testing.T){
    lru := New(0, nil) 
    lru.Add("k1", String("v1"))
    lru.Add("k2", String("v2"))
    lru.Add("k3", String("v3"))
    lru.Add("k4", String("v4"))
    lru.Add("k1", String("1v"))
    if lru.Len() != 4{
        t.Error("Add: Wrong c.Len()")
    }
    if v1, ok := lru.Get("k1"); !ok || v1 != String("1v") {
        t.Error("Add: Wrong update")
    }
}

func TestRemoveOldest(t * testing.T){
    k1, k2, k3 := "k1", "k2", "k3"
    v1, v2, v3 := "v1", "v2", "v3"
    cap := uint64(len(k1 + k2 + v1 + v2))
    lru := New(cap, nil) 
    lru.Add(k1, String(v1))
    lru.Add(k2, String(v2))
    lru.Add(k3, String(v3))
    if _, ok := lru.Get("k1"); ok || lru.Len() != 2 {
        t.Error("RemoveOldest failed")
    }

}

// func TestMain(m * testing.M){
//
// }

