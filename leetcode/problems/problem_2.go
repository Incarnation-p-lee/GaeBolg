package problems

import "fmt"

func appendTail(last *ListNode, l *ListNode, carry bool) {
	for ; l != nil; l = l.Next {
		v := l.Val

		if carry {
			carry = false
			v++
		}

		if v > 9 {
			carry = true
			v -= 10
		}

		node := new(ListNode)
		node.Val = v
		last.Next = node
		last = node
	}

	if carry {
		node := new(ListNode)
		node.Val = 1
		last.Next = node
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry bool
	var out ListNode
	var last *ListNode

	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1 == nil && l2 == nil {
		return nil
	}

	last = &out

	for ; l1 != nil && l2 != nil; l1 = l1.Next {
		v := l1.Val + l2.Val

		if carry {
			carry = false
			v++
		}

		if v > 9 {
			carry = true
			v -= 10
		}

		node := new(ListNode)
		node.Val = v
		last.Next = node
		last = node

		l2 = l2.Next
	}

	if l1 != nil {
		appendTail(last, l1, carry)
	} else if l2 != nil {
		appendTail(last, l2, carry)
	} else if carry {
		node := new(ListNode)
		node.Val = 1
		last.Next = node
	}

	return out.Next;
}

func AddTwoNumbers() {
	var l1, l2, l3 ListNode

	l1.Val = 9
	l2.Val = 8
	l3.Val = 1

	l1.Next = &l2
	l2.Next = nil
	l3.Next = nil

	fmt.Printf("<002> [")

	for l := addTwoNumbers(&l1, &l3); l != nil; l = l.Next {
		fmt.Printf("%d, ", l.Val)
	}

	fmt.Printf("]\n")
}

