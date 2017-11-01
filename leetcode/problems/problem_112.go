package problems

import "fmt"

func hasPathSumDfs(root *TreeNode, sum int) bool {
	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			return true
		} else {
			return false
		}
	} else {
		if root.Left != nil {
			if hasPathSumDfs(root.Left, sum - root.Val) {
				return true
			}
		}

		if root.Right != nil {
			if hasPathSumDfs(root.Right, sum - root.Val) {
				return true
			}
		}

		return false
	}
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	return hasPathSumDfs(root, sum)
}

func HasPathSum() {
	var t1, t2, t3, t4, t5, t6, t7 TreeNode

	t1.Val, t2.Val, t3.Val = 2, 4, 6
	t2.Left, t2.Right = &t1, &t3

	t4.Val, t5.Val, t6.Val, t7.Val = 1, 3, 5, 7
	t1.Left, t1.Right = &t4, &t5
	t3.Left, t3.Right = &t6, &t7
	t4.Left, t4.Right, t5.Left, t5.Right = nil, nil, nil, nil
	t6.Left, t6.Right, t7.Left, t7.Right = nil, nil, nil, nil

	fmt.Printf("<112> ")
	fmt.Println(hasPathSum(&t2, 15))
}

