package problems

import "fmt"

func isSameTreeOrder(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreeOrder(p.Left, q.Left) && isSameTreeOrder(p.Right, q.Right)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	return isSameTreeOrder(p, q)
}

func IsSameTree() {
	var t1, t2, t3, t4, t5, t6 TreeNode

	t1.Val, t2.Val, t3.Val = 2, 1, 3
	t2.Left, t2.Right = &t1, &t3
	t1.Left, t1.Right, t3.Left, t3.Right = nil, nil, nil, nil

	t4.Val, t5.Val, t6.Val = 2, 1, 6
	t5.Left, t5.Right = &t4, &t6
	t4.Left, t4.Right, t6.Left, t6.Right = nil, nil, nil, nil

	fmt.Printf("<100> ")
	fmt.Println(isSameTree(&t2, &t5))
}

