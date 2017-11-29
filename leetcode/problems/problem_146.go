package problems

import "fmt"

type DListNode struct {
	Val int
	Prev, Next *DListNode
}

type LRUCache struct {
	lmap map[int]*DListNode
	Head DListNode
}

func Constructor(capacity int) LRUCache {
	var cache LRUCache

	cache.lmap = make(map[int]*DListNode)
	h := &cache.Head

	for i := 0; i < capacity; i++ {
		node := new(DListNode)
		h.Next = node
		node.Prev = h
		h = h.Next
	}

	h.Next = &cache.Head
	cache.Head.Prev = h

	return cache
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
	insertToFirst(&this.Head, node)

	fmt.Println(this.lmap)

	return v.Val
}


func (this *LRUCache) Put(key int, value int)  {
	if this == nil {
		return
	} 

	_, ok := this.lmap[key]

	last := this.Head.Prev
	last.Val = value
	insertToFirst(&this.Head, last)

	if !ok {
		this.lmap[key] = last
	}
}

func LRUCacheSample() {
	cache := Constructor(2)
	fmt.Printf("<146> ")

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Printf("%d ", cache.Get(1))

	cache.Put(3, 3)
	fmt.Printf("%d ", cache.Get(2))

	cache.Put(4, 4)
	fmt.Printf("%d ", cache.Get(1))
	fmt.Printf("%d ", cache.Get(3))
	fmt.Printf("%d ", cache.Get(4))

	fmt.Println()
}

