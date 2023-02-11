package dsa

import (
	"container/list"
	"sync"
	"unsafe"
)

type LRUCache struct {
	cap int

	list *list.List

	items map[int]*list.Element // 记录链表中有哪些元素，这样就不用遍历链表了，

	mu sync.Mutex
}

type Value struct {
	key   int
	value int
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		cap:   cap,
		list:  list.New(),
		items: make(map[int]*list.Element),
	}
}

func (c *LRUCache) Get(key int) int {

	c.mu.Lock()
	defer c.mu.Unlock()

	if item, exit := c.items[key]; exit {
		c.list.MoveToFront(item)
		return (*Value)(unsafe.Pointer(&item.Value)).value
	}
	return -1
}

func (c *LRUCache) Put(key, val int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if item, exit := c.items[key]; exit {
		c.list.MoveToFront(item)
		(*Value)(unsafe.Pointer(&item.Value)).value = val
	} else {
		e := c.list.PushFront(&Value{key, val})
		c.items[key] = e
		if (c.list.Len()) > c.cap {
			// 删除链表 末尾 也是最久没用过 的元素
			c.removeOldest()

		}
	}

}

func (c *LRUCache) removeOldest() {
	e := c.list.Back()
	if e != nil {
		c.list.Remove(e)
		v := (*Value)(unsafe.Pointer(&e.Value))
		delete(c.items, v.key)
	}
}
