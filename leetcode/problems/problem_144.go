package problems

import "fmt"

func preorderTraversal(root *TreeNode) []int {
	out := make([]int, 0)

	var preOrderRecursive func (node *TreeNode)
	preOrderRecursive = func (node *TreeNode) {
		if node != nil {
			out = append(out, node.Val)

			preOrderRecursive(node.Left)
			preOrderRecursive(node.Right)
		}
	}

	preOrderRecursive(root)

	return out
}

func PreorderTraversal() {
	var t1, t2, t3 TreeNode

	t1.Val, t2.Val, t3.Val = 2, 1, 3
	t2.Left, t2.Right = &t1, &t3
	t1.Left, t1.Right, t3.Left, t3.Right = nil, nil, nil, nil

	fmt.Printf("<144> ")
	fmt.Println(preorderTraversal(&t2))
}

