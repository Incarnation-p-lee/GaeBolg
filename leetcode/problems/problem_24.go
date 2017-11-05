package problems

import "fmt"

func swapPairs(head *ListNode) *ListNode {
	var faker ListNode

	if head == nil || head.Next == nil {
		return head
	}

	p1, p2, h := head, head.Next, &faker

	for p1 != nil && p2 != nil {
		next := p2.Next

		h.Next = p2
		p2.Next = p1
		p1.Next = next

		if next == nil || next.Next == nil {
			break
		}

		h = p1
		p1, p2 = next, next.Next
	}


	return faker.Next
}

func SwapPairs() {
	var l1, l2, l3, l4, l5, l6, l7 ListNode

	l1.Val, l2.Val, l3.Val, l4.Val, l5.Val, l6.Val, l7.Val = 1, 2, 3, 4, 5, 6, 7

	l1.Next, l2.Next, l3.Next, l4.Next, l5.Next = &l2, &l3, &l4, &l5, &l6
	l6.Next, l7.Next = &l7, nil

	fmt.Printf("<024> ")
	fmt.Println(swapPairs(&l1))
}

