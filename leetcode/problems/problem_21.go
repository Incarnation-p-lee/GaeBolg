package problems

import "fmt"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var faker ListNode

	head = &faker

	for l1 != nil && l2 != nil {
		node := new(ListNode)

		if l1.Val < l2.Val {
			node.Val = l1.Val
			l1 = l1.Next
		} else {
			node.Val = l2.Val
			l2 = l2.Next
		}

		head.Next = node
		head = node
	}

	var listAppend func (n *ListNode)
	listAppend = func (n *ListNode) {
		for n != nil {
			node := new(ListNode)
			node.Val = n.Val

			head.Next = node
			head = node
			n = n.Next
		}
	}

	if l1 != nil {
		listAppend(l1)
	}

	if l2 != nil {
		listAppend(l2)
	}

	return faker.Next
}

func MergeTwoLists() {
	var l1, l2, l3, l4, l5, l6, l7 ListNode

	l1.Val, l2.Val, l3.Val, l4.Val, l5.Val, l6.Val, l7.Val = 1, 2, 3, 4, 5, 6, 7
	l1.Next, l2.Next, l3.Next, l4.Next = &l2, &l3, &l4, nil
	l5.Next, l6.Next, l7.Next = &l6, &l7, nil

	fmt.Printf("<021> ")
	fmt.Println(mergeTwoLists(&l1, &l2))
}

