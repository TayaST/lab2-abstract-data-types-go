package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	cap int
	items map[int]*list.Element
	lst *list.List
}

type cacheItem struct { key, value int }

func (c *LRUCache) Init(capacity int) {
	c.cap = capacity
	c.items = make(map[int]*list.Element)
	c.lst = list.New()
}

func (c *LRUCache) Set(key, value int) {
	if elem, ok := c.items[key]; ok {
		c.lst.MoveToFront(elem)
		elem.Value.(*cacheItem).value = value
		return
	}
	if c.lst.Len() >= c.cap {
		back := c.lst.Back()
		delete(c.items, back.Value.(*cacheItem).key)
		c.lst.Remove(back)
	}
	elem := c.lst.PushFront(&cacheItem{key, value})
	c.items[key] = elem
}

func (c *LRUCache) Get(key int) int {
	if elem, ok := c.items[key]; ok {
		c.lst.MoveToFront(elem)
		return elem.Value.(*cacheItem).value
	}
	return -1
}

func runCache() {
	fmt.Println("\n=== Задание 7: LRU Кэш ===")
	var cap, q int
	fmt.Print("Введите cap и Q: "); fmt.Scan(&cap, &q)
	c := &LRUCache{}
	c.Init(cap)

	for i := 0; i < q; i++ {
		var cmd string
		fmt.Scan(&cmd)
		if cmd == "SET" {
			var x, y int
			fmt.Scan(&x, &y)
			c.Set(x, y)
		} else if cmd == "GET" {
			var x int
			fmt.Scan(&x)
			fmt.Print(c.Get(x), " ")
		}
	}
	fmt.Println()
}