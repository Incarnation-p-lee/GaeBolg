package problems

import "fmt"
import "math"

var max int

func maxPathSumOrder(n *TreeNode) int {
	var d int

	if n == nil {
		return 0
	}

	left := maxPathSumOrder(n.Left)
	right := maxPathSumOrder(n.Right)

	d = n.Val

	if left > 0 {
		d += left
	}

	if right > 0 {
		d += right
	}

	if d > max {
		max = d
	}

	var maxInt func(a, b int) int
	maxInt = func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	return maxInt(n.Val, maxInt(n.Val + left, n.Val + right))
}

func maxPathSum(root *TreeNode) int {
	max = math.MinInt32

	maxPathSumOrder(root)

	return max
}

func MaxPathSum() {
	var t1, t2, t3, t4, t5, t6, t7 TreeNode

	t1.Val, t2.Val, t3.Val = 2, 4, 6
	t2.Left, t2.Right = &t1, &t3

	t4.Val, t5.Val, t6.Val, t7.Val = 1, 3, 5, 7
	t1.Left, t1.Right = &t4, &t5
	t3.Left, t3.Right = &t6, &t7
	t4.Left, t4.Right, t5.Left, t5.Right = nil, nil, nil, nil
	t6.Left, t6.Right, t7.Left, t7.Right = nil, nil, nil, nil

	fmt.Printf("<124> ")
	fmt.Println(maxPathSum(&t2))
}

