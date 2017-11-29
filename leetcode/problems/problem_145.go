package problems

import "fmt"

func postorderTraversal(root *TreeNode) []int {
	out := make([]int, 0)

	var postOrderRecursive func (node *TreeNode)
	postOrderRecursive = func (node *TreeNode) {
		if node != nil {
			postOrderRecursive(node.Left)
			postOrderRecursive(node.Right)

			out = append(out, node.Val)
		}
	}

	postOrderRecursive(root)

	return out
}

func PostorderTraversal() {
	var t1, t2, t3 TreeNode

	t1.Val, t2.Val, t3.Val = 2, 1, 3
	t2.Left, t2.Right = &t1, &t3
	t1.Left, t1.Right, t3.Left, t3.Right = nil, nil, nil, nil

	fmt.Printf("<145> ")
	fmt.Println(postorderTraversal(&t2))
}

