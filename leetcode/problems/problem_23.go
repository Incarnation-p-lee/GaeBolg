package problems

import "fmt"

type minHeap struct {
	Size int
	Array []*ListNode
}

func minHeapCreate(size int) * minHeap {
	var heap *minHeap

	heap = new(minHeap)

	heap.Size = 0
	heap.Array = make([]*ListNode, size + 1)

	return heap
}

func minHeapRootRemove(heap *minHeap) *ListNode {
	array := heap.Array
	root := array[1]
	last := array[heap.Size]
	array[heap.Size] = nil
	heap.Size -= 1

	i := 1

	var minNode func (j int) int
	minNode = func (j int) int {
		if array[j * 2].Val > array[j * 2 + 1].Val {
			return j * 2 + 1
		} else {
			return j * 2
		}
	}

	for {
		if i * 2 > heap.Size {
			array[i] = last
			break
		} else if i * 2 == heap.Size {
			if last.Val > array[i * 2].Val {
				array[i], array[i * 2] = array[i * 2], last
			} else {
				array[i] = last
			}
			break
		} else {
			k := minNode(i)
			if last.Val > array[k].Val {
				array[i], array[k] = array[k], last
				i = k
			} else {
				array[i] = last
				break
			}
		}
	}

	return root
}

func minHeapNodeInsert(heap *minHeap, inserted *ListNode) {
	array := heap.Array
	heap.Size += 1

	i := heap.Size
	array[i] = inserted

	for {
		if i == 1 || array[i / 2].Val <= inserted.Val {
			array[i] = inserted
			break
		} else {
			array[i] = array[i / 2]
			i = i / 2
		}
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	var faker ListNode
	var head *ListNode
	var heap *minHeap

	head = &faker 
	heap = minHeapCreate(len(lists))

	for _, list := range lists {
		if list != nil {
			minHeapNodeInsert(heap, list)
		}
	}

	for heap.Size != 0 {
		min := minHeapRootRemove(heap)
		node := new(ListNode)
		node.Val = min.Val

		head.Next = node
		head = node

		if min.Next != nil {
			min = min.Next
			minHeapNodeInsert(heap, min)
		}
	}

	return faker.Next
}

func MergeKLists() {
	var l1, l2, l3, l4, l5, l6, l7 ListNode

	l1.Val, l2.Val, l3.Val, l4.Val, l5.Val, l6.Val, l7.Val = 2, 3, 4, 1, 9, 7, 6
	l1.Next, l2.Next, l3.Next = nil, nil, nil
	l4.Next, l5.Next, l6.Next = nil, nil, nil
	l7.Next = nil

	fmt.Printf("<023> ")
	fmt.Println(mergeKLists([]*ListNode{&l1, &l2, &l3, &l4, &l5, &l6, &l7}))
}

