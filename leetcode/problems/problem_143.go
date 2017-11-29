package problems

import "fmt"

func reorderList(head *ListNode) {
	h := head

	if h == nil {
		return
	}

	faker1 := new(ListNode)
	faker2 := new(ListNode)

	p1, p2 := h, h

	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next 
	}

	limit := p1

	for p1 != nil {
		next, p1Next := faker1.Next, p1.Next
		faker1.Next = p1
		p1.Next = next

		p1 = p1Next
	}

	p1, p2 = h, faker1.Next
	tmp := faker2

	for p1 != limit {
		p1Next, p2Next := p1.Next, p2.Next
		tmp.Next = p1
		tmp = tmp.Next
		tmp.Next = p2
		tmp = tmp.Next

		p1, p2 = p1Next, p2Next
	}

	if p2 != nil {
		tmp.Next = p2
	}
}

func ReorderList() {
	var l1, l2, l3, l4, l5, l6, l7 ListNode

	l1.Val, l2.Val, l3.Val, l4.Val, l5.Val, l6.Val, l7.Val = 1, 2, 3, 4, 5, 6, 7
	l1.Next, l2.Next, l3.Next, l4.Next = &l2, &l3, &l4, &l5
	l5.Next, l6.Next, l7.Next = &l6, &l7, nil

	fmt.Printf("<143> ")
    reorderList(&l1)
	fmt.Println(l1)
}

