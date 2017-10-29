package problems

import "fmt"

func buildTreeOrder(pOrder []int, iOrder []int) *TreeNode {
	var node *TreeNode

	if len(pOrder) == 0 || len(iOrder) == 0 {
		return nil
	}

	node = new(TreeNode)
	node.Val = pOrder[0]

	idx := func (t int) int {
		for i, v := range iOrder {
			if v == t {
				return i
			}
		}
		return 0
	}(pOrder[0])

	node.Left = buildTreeOrder(pOrder[1:idx + 1], iOrder[0:idx])
	node.Right = buildTreeOrder(pOrder[idx + 1:], iOrder[idx + 1:])

	return node
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) {
		return nil
	} else if len(preorder) == 0 {
		return nil
	}

	return buildTreeOrder(preorder, inorder)
}

func BuildTree() {
	pOrder := []int{4, 2, 1, 3, 6, 5}
	iOrder := []int{1, 2, 3, 4, 5, 6}

	fmt.Printf("<105> ")
	fmt.Println(buildTree(pOrder, iOrder))
}

