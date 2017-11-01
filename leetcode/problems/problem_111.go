package problems

import "fmt"

func minDepthDfs(node *TreeNode, deep int, m *int) {
	if node.Left == nil && node.Right == nil {
		if deep < *m {
			*m = deep
		}
	} else {
		if node.Left != nil {
			minDepthDfs(node.Left, deep + 1, m)
		}

		if node.Right != nil {
			minDepthDfs(node.Right, deep + 1, m)
		}
	}
}

func minDepth(root *TreeNode) int {
	var m int = 0x7fffffff

	if root == nil {
		return 0
	}

	minDepthDfs(root, 0, &m)

	return m
}

func MinDepth() {
	var t1, t2, t3, t4, t5, t6 TreeNode

	t1.Val, t2.Val, t3.Val = 1, 2, 3
	t4.Val, t5.Val, t6.Val = 5, 4, 6

	t2.Left, t2.Right = &t1, nil
	t1.Left, t1.Right, t3.Left, t3.Right = &t3, nil, &t4, nil

	t4.Left, t4.Right = &t5, &t6
	t5.Left, t5.Right, t6.Left, t6.Right = nil, nil, nil, nil

	fmt.Printf("<111> ") 
	fmt.Println(minDepth(&t2))
} 

