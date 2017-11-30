package problems

import "fmt"

func sortList(head *ListNode) *ListNode {
	h := head

	if h == nil {
		return h
	}

	var mergeList func (l1, l2 *ListNode) *ListNode
	mergeList = func (l1, l2 *ListNode) *ListNode {
		faker := new(ListNode)
		l := faker

		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				l.Next = l1
				l1 = l1.Next
			} else {
				l.Next = l2
				l2 = l2.Next
			}

			l = l.Next
		}

		if l1 != nil {
			l.Next = l1
		}

		if l2 != nil {
			l.Next = l2
		}

		return faker.Next
	}

	var mergeSort func (list *ListNode) *ListNode
	mergeSort = func (list *ListNode) *ListNode {
		if list == nil {
			return nil
		} else if list.Next == nil {
			return list
		}

		p1, p2, p1Prev := list, list, list

		for p2 != nil && p2.Next != nil {
			p1Prev = p1
			p1 = p1.Next
			p2 = p2.Next.Next
		}

		left, right := list, p1
		p1Prev.Next = nil

		p1 = mergeSort(left)
		p2 = mergeSort(right)

		return mergeList(p1, p2)
	}

	return mergeSort(h)
}

func SortList() {
	var l1, l2, l3, l4, l5, l6, l7 ListNode

	l1.Val, l2.Val, l3.Val, l4.Val, l5.Val, l6.Val, l7.Val = 7, 5, 3, 4, 2, 6, 1
	l1.Next, l2.Next, l3.Next, l4.Next = &l2, &l3, &l4, &l5
	l5.Next, l6.Next, l7.Next = &l6, &l7, nil

	fmt.Printf("<148> ")
	fmt.Println(sortList(&l1))
}

