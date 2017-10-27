package problems

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isValidBSTInOrder(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if isValidBSTInOrder(root.Left) == false {
		return false
	}

	if last != root && last.Val >= root.Val {
		return false
	}

	last = root

	if isValidBSTInOrder(root.Right) == false {
		return false
	}

	return true
}

var last *TreeNode

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	last = root

	for ; last.Left != nil; last = last.Left {
		;
	}

	return isValidBSTInOrder(root)
}

func IsValidBST() {
	var t1, t2, t3, t4, t5, t6 TreeNode

	t1.Val, t2.Val, t3.Val = 1, 2, 3
	t2.Left, t2.Right = &t1, &t3
	t1.Left, t1.Right, t3.Left, t3.Right = nil, nil, nil, nil

	t4.Val, t5.Val, t6.Val = 5, 4, 6
	t5.Left, t5.Right = &t4, &t6
	t4.Left, t4.Right, t6.Left, t6.Right = nil, nil, nil, nil

	fmt.Printf("<098> ")
	fmt.Println(isValidBST(&t2))

	fmt.Printf("<098> ")
	fmt.Println(isValidBST(&t5))
}

