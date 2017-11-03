package problems

import "fmt"

func sumNumbersDfs(n *TreeNode, val int, sum *int) {
	if n == nil {
		return
	} else if n.Left == nil && n.Right == nil {
		*sum += val * 10 + n.Val
		return
	}

	sumNumbersDfs(n.Left, val * 10 + n.Val, sum)
	sumNumbersDfs(n.Right, val * 10 + n.Val, sum)
}

func sumNumbers(root *TreeNode) int {
	var sum int

	if root == nil {
		return 0
	}

	sumNumbersDfs(root, 0, &sum)

	return sum
}

func SumNumbers() {
	var t1, t2, t3 TreeNode

	t1.Val, t2.Val, t3.Val = 2, 1, 3

	t2.Left, t2.Right = &t1, &t3
	t1.Left, t1.Right, t3.Left, t3.Right = nil, nil, nil, nil

	fmt.Printf("<129> ")
	fmt.Println(sumNumbers(&t2))
}

