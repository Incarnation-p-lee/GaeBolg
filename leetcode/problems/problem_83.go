package problems

import "fmt"

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var faker ListNode

	n := &faker
	n.Next = head
	n = n.Next

	last := head.Val

	for l := head.Next; l != nil; l = l.Next {
		if l.Val != last {
			last = l.Val
			n.Next = l
			n = n.Next
		}
	}

	n.Next = nil

	return faker.Next
}

func DeleteDuplicates() {
	var l1, l2, l3, l4, l5, l6, l7 ListNode

	l1.Val, l2.Val, l3.Val, l4.Val, l5.Val, l6.Val, l7.Val = 1, 1, 3, 4, 5, 5, 7
	l1.Next, l2.Next, l3.Next, l4.Next = &l2, &l3, &l4, &l5
	l5.Next, l6.Next, l7.Next = &l6, &l7, nil

	fmt.Printf("<083> ")
	fmt.Println(deleteDuplicates(&l1))
}

