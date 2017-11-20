package problems

import "fmt"

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	var faker ListNode
	n := &faker

	for l := head; l != nil; l = l.Next {
		count, val := 1, l.Val

		for l.Next != nil && l.Next.Val == val {
			count++
			l = l.Next
		}

		if count == 1 {
			n.Next = l
			n = n.Next
		}
	}

	n.Next = nil

	return faker.Next
}

func DeleteDuplicates() {
	var l1, l2, l3, l4, l5, l6, l7 ListNode

	l1.Val, l2.Val, l3.Val, l4.Val, l5.Val, l6.Val, l7.Val = 1, 1, 2, 3, 3, 3, 3
	l1.Next, l2.Next, l3.Next, l4.Next = &l2, &l3, &l4, &l5
	l5.Next, l6.Next, l7.Next = &l6, &l7, nil

	fmt.Printf("<082> ")
	fmt.Println(deleteDuplicates(&l1))
}

