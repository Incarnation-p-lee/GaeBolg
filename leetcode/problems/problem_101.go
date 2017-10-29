package problems

import "fmt"

func isSymmetricOrder(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil {
		return false
	} else if q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSymmetricOrder(p.Left, q.Right) && isSymmetricOrder(p.Right, q.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricOrder(root.Left, root.Right)
}

func IsSymmetric() {
	var t1, t2, t3, t4, t5, t6, t7 TreeNode

	t1.Val, t2.Val, t3.Val = 2, 1, 2
	t2.Left, t2.Right = &t1, &t3

	t4.Val, t5.Val, t6.Val, t7.Val = 4, 5, 5, 4
	t1.Left, t1.Right = &t4, &t5
	t3.Left, t3.Right = &t6, &t7
	t4.Left, t4.Right, t5.Left, t5.Right = nil, nil, nil, nil
	t6.Left, t6.Right, t7.Left, t7.Right = nil, nil, nil, nil

	fmt.Printf("<101> ")
	fmt.Println(isSymmetric(&t2))
}

