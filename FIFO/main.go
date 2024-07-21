package main

import (
	"container/list"
	"fmt"
)

type FIFOCache struct {
	maxEntries int
	ll         *list.List
	cache      map[interface{}]*list.Element
}

type entry struct {
	key   interface{}
	value interface{}
}

func NewFIFO(maxEntries int) *FIFOCache {
	return &FIFOCache{
		maxEntries: maxEntries,
		ll:         list.New(),
		cache:      make(map[interface{}]*list.Element),
	}
}

func (c *FIFOCache) Get(key interface{}) (value interface{}, ok bool) {
	if ele, hit := c.cache[key]; hit {
		return ele.Value.(*entry).value, true
	}
	return
}

func (c *FIFOCache) Add(key, value interface{}) {
	if ele, hit := c.cache[key]; hit {
		ele.Value.(*entry).value = value
		return
	}
	ele := c.ll.PushBack(&entry{key, value})
	c.cache[key] = ele
	if c.maxEntries != 0 && c.ll.Len() > c.maxEntries {
		c.RemoveOldest()
	}
}

func (c *FIFOCache) RemoveOldest() {
	ele := c.ll.Front()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
	}
}

func (c *FIFOCache) Len() int {
	return c.ll.Len()
}

func main() {
	c := NewFIFO(2)
	c.Add("a", 1)
	c.Add("b", 2)
	fmt.Println(c.Get("a"))
	c.Add("c", 3)
	fmt.Println(c.Get("a"))
	fmt.Println(c.Get("b"))
	fmt.Println(c.Get("c"))
}
