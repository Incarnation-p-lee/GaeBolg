package problems

import "fmt"

func inorderTraversal(root *TreeNode) []int {
	n := root

	if n == nil {
		return []int{}
	}

	out := make([]int, 0)

	var traversalInorder func (node *TreeNode)
	traversalInorder = func (node *TreeNode) {
		if node != nil {
			traversalInorder(node.Left)
			out = append(out, node.Val)
			traversalInorder(node.Right)
		}
	}

	traversalInorder(n)

	return out
}

func InorderTraversal() {
	var t1, t2, t3, t4, t5, t6 TreeNode

	t1.Val, t2.Val, t3.Val = 3, 2, 1
	t1.Left, t1.Right, t2.Left = &t2, &t3, &t5
	t2.Right, t3.Left, t3.Right = nil, nil, nil

	t4.Val, t5.Val, t6.Val = 5, 4, 6
	t5.Left, t5.Right = &t4, &t6
	t4.Left, t4.Right, t6.Left, t6.Right = nil, nil, nil, nil

	fmt.Printf("<094> ")
	fmt.Println(inorderTraversal(&t1))
}

