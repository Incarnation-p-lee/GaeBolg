package problems

import "fmt"

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	h := head

	if h == nil {
		return nil
	}

	k := 1
	faker, ext := new(ListNode), new(ListNode)
	ls := faker

	for ; h != nil; h = h.Next {
		if k == m {
			node, last := h, h

			for k <= n {
				tmp := node.Next
				node.Next = ext.Next
				ext.Next = node

				k++
				node = tmp
			}

			last.Next = node
			ls.Next = ext.Next

			break
		}

		k++
		ls.Next = h
		ls = ls.Next
	}

	return faker.Next
}

func ReverseBetween() {
	var l1, l2, l3, l4, l5, l6, l7 ListNode

	l1.Val, l2.Val, l3.Val, l4.Val, l5.Val, l6.Val, l7.Val = 1, 2, 3, 4, 5, 6, 7
	l1.Next, l2.Next, l3.Next, l4.Next = &l2, &l3, &l4, &l5
	l5.Next, l6.Next, l7.Next = &l6, &l7, nil

	fmt.Printf("<092> ")
	fmt.Println(reverseBetween(&l1, 2, 5))
}

