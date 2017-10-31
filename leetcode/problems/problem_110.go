package problems

import "fmt"
import "math"

func treeHeight(t *TreeNode) int {
	if t == nil {
		return 0
	}

	l := treeHeight(t.Left)
	r := treeHeight(t.Right)

	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}    

	if !isBalanced(root.Left) {
		return false
	}

	if !isBalanced(root.Right) {
		return false
	}

	if math.Abs(float64(treeHeight(root.Left) - treeHeight(root.Right))) > 1 {
		return false
	}

	return true
}

func IsBalanced() {
	var t1, t2, t3, t4, t5, t6 TreeNode

	t1.Val, t2.Val, t3.Val = 1, 2, 3
	t2.Left, t2.Right = &t1, &t3
	t1.Left, t1.Right, t3.Left, t3.Right = nil, nil, nil, nil

	t4.Val, t5.Val, t6.Val = 5, 4, 6
	t5.Left, t5.Right = &t4, &t6
	t4.Left, t4.Right, t6.Left, t6.Right = nil, nil, nil, nil

	fmt.Printf("<110> ") 
	fmt.Println(isBalanced(&t2))
}

