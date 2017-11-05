package problems

import "fmt"

func reverseKGroup(head *ListNode, k int) *ListNode {
	var faker ListNode
	var stack []*ListNode

	if head == nil || k < 2 {
		return head
	}

	prev := &faker
	prev.Next = head

	for head != nil {
		if len(stack) < k {
			stack = append(stack, head)
			head = head.Next
		}

		if len(stack) == k {
			p1 := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]

			next := p1.Next
			prev.Next = p1

			for len(stack) > 0 {
				p0 := stack[len(stack) - 1]
				stack = stack[:len(stack) - 1]

				p1.Next, p1 = p0, p0
			}

			p1.Next = next
			prev = p1
		}
	}

	return faker.Next
}

func ReverseKGroup() {
	var l1, l2, l3, l4, l5, l6, l7 ListNode

	l1.Val, l2.Val, l3.Val, l4.Val, l5.Val, l6.Val, l7.Val = 1, 2, 3, 4, 5, 6, 7

	l1.Next, l2.Next, l3.Next, l4.Next, l5.Next = &l2, &l3, &l4, &l5, &l6
	l6.Next, l7.Next = &l7, nil

	fmt.Printf("<025> ")
	fmt.Println(reverseKGroup(&l1, 4))
}

