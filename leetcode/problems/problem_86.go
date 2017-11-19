package problems

import "fmt"

func partitionList(head *ListNode, x int) *ListNode {
	h := head

	if h == nil {
		return nil
	}

	small, big := new(ListNode), new(ListNode)

	s, b := small, big

	for h != nil {
		if h.Val < x {
			s.Next = h
			s = s.Next
		} else {
			b.Next = h
			b = b.Next
		}

		h = h.Next
	}

	s.Next, b.Next = big.Next, nil

	return small.Next
}

func PartitionList() {
	var l1, l2, l3, l4, l5, l6, l7 ListNode

	l1.Val, l2.Val, l3.Val, l4.Val, l5.Val, l6.Val, l7.Val = 7, 2, 6, 4, 1, 5, 3
	l1.Next, l2.Next, l3.Next, l4.Next = &l2, &l3, &l4, &l5
	l5.Next, l6.Next, l7.Next = &l6, &l7, nil

	fmt.Printf("<086> ")
	fmt.Println(partitionList(&l1, 3))
}

