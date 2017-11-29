package problems

import "fmt"

type DListNode struct {
	Val, Key int
	Prev, Next *DListNode
}

type LRUCache struct {
	Head *DListNode
	lmap map[int]*DListNode
}

func Constructor(capacity int) LRUCache {
	cache := new(LRUCache)

	cache.lmap = make(map[int]*DListNode)
	cache.Head = new(DListNode)
	cache.Head.Prev, cache.Head.Next = cache.Head, cache.Head
	h := cache.Head

	for i := 0; i < capacity; i++ {
		node := new(DListNode)

		node.Next = h.Next
		node.Prev = h

		h.Next.Prev = node
		h.Next = node
	}

	return *cache
}

func insertToFirst(head, node *DListNode) {
	if head == nil || node == nil {
		return
	}

	if node.Prev == head {
		return
	}

	prev := node.Prev
	next := node.Next

	prev.Next = next
	next.Prev = prev

	node.Prev = head
	node.Next = head.Next
	head.Next.Prev = node
	head.Next = node
}

func (this *LRUCache) Get(key int) int {
	if this == nil {
		return -1
	} 

	v, ok := this.lmap[key]

	if !ok {
		return -1
	}

	node := v
	insertToFirst(this.Head, node)

	return v.Val
}

func (this *LRUCache) Put(key int, value int)  {
	if this == nil {
		return
	} 

	node, ok := this.lmap[key]

	if ok {
		delete(this.lmap, key)

		node.Key, node.Val = key, value
		insertToFirst(this.Head, node)

		this.lmap[key] = node
	} else {
		last := this.Head.Prev
		oldKey := last.Key

		_, ok := this.lmap[oldKey]

		if ok {
			delete(this.lmap, oldKey)
		}

		last.Key, last.Val = key, value
		insertToFirst(this.Head, last)

		this.lmap[key] = last
	}
}

func LRUCacheSample() {
	cache := Constructor(2)
	fmt.Printf("<146> ")

	fmt.Printf("%d ", cache.Get(2))
	cache.Put(2, 6)
	fmt.Printf("%d ", cache.Get(1))

	cache.Put(1, 5)
	cache.Put(1, 2)
	fmt.Printf("%d ", cache.Get(1))
	fmt.Printf("%d ", cache.Get(2))

	fmt.Println()
}

