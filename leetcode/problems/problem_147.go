package problems

import "fmt"

func insertionSortList(head *ListNode) *ListNode {
	h := head

	if h == nil {
		return nil
	} else if h.Next == nil {
		return h
	}

	faker := new(ListNode)
	faker.Next = h
	h = h.Next
	faker.Next.Next = nil

	for h != nil {
		prev, l := faker, faker.Next

		for l != nil {
			if l.Val > h.Val {
				break
			}

			prev = l
			l = l.Next
		}

		next := h.Next

		prev.Next = h
		h.Next = l

		h = next
	}

	return faker.Next
}

func InsertionSortList() {
	var l1, l2, l3, l4, l5, l6, l7 ListNode

	l1.Val, l2.Val, l3.Val, l4.Val, l5.Val, l6.Val, l7.Val = 4, 7, 5, 1, 3, 6, 2
	l1.Next, l2.Next, l3.Next, l4.Next = &l2, &l3, &l4, &l5
	l5.Next, l6.Next, l7.Next = &l6, &l7, nil

	fmt.Printf("<147> ")
	fmt.Println(insertionSortList(&l1))
}

