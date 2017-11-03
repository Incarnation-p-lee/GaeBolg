package problems

import "fmt"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var prev, node *ListNode

	if head == nil {
		return nil
	}

	d, node := n, head
	for node != nil {
		if prev == nil && d != 0 {
			d--
		} else if prev == nil && d == 0 {
			prev = head
		} else if prev != nil {
			prev = prev.Next
		}

		node = node.Next
	}

	if prev == nil {
		return head.Next
	}

	target := prev.Next
	prev.Next = target.Next

	return head
}

func RemoveNthFromEnd() {
	var l1, l2, l3, l4, l5, l6, l7 ListNode

	l1.Val, l2.Val, l3.Val, l4.Val, l5.Val, l6.Val, l7.Val = 1, 2, 3, 4, 5, 6, 7
	l1.Next, l2.Next, l3.Next, l4.Next = &l2, &l3, &l4, &l5
	l5.Next, l6.Next, l7.Next = &l6, &l7, nil

	fmt.Printf("<019> ")
	fmt.Println(removeNthFromEnd(&l1, 2))
}

