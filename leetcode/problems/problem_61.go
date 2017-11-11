package problems

import "fmt"

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	l := 0

	for ls := head; ls != nil; ls = ls.Next {
		l++
	}

	k = k % l

	if k == 0 {
		return head
	}

	p2 := head

	for k > 0 {
		p2 = p2.Next
		k--
	}

	p1 := head

	for p2.Next != nil {
		p2 = p2.Next
		p1 = p1.Next
	}

	out := p1.Next
	p1.Next, p2.Next = nil, head

	return out
}

func RotateRight() {
	var l1, l2, l3, l4, l5, l6, l7 ListNode

	l1.Val, l2.Val, l3.Val, l4.Val, l5.Val, l6.Val, l7.Val = 1, 2, 3, 4, 5, 6, 7
	l1.Next, l2.Next, l3.Next, l4.Next = &l2, &l3, &l4, &l5
	l5.Next, l6.Next, l7.Next = &l6, &l7, nil

	fmt.Printf("<061> ")
	fmt.Println(rotateRight(&l1, 3))
}

