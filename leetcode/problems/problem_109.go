package problems

import "fmt"

func sortedListToBST(head *ListNode) *TreeNode {
	var data []int

	for ; head != nil; head = head.Next {
		data = append(data, head.Val)
	}

	return sortedArrayToBST(data)
}

func SortedListToBST() {
	var l1, l2, l3, l4, l5, l6, l7 ListNode

	l1.Val, l2.Val, l3.Val, l4.Val, l5.Val, l6.Val, l7.Val = 1, 2, 3, 4, 5, 6, 7
	l1.Next, l2.Next, l3.Next, l4.Next = &l2, &l3, &l4, &l5
	l5.Next, l6.Next, l7.Next = &l6, &l7, nil

	fmt.Printf("<109> ")
	fmt.Println(sortedListToBST(&l1))
}

