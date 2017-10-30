package problems

import "fmt"

func buildTree2Order(iOrder []int, pOrder []int) *TreeNode {
	var node *TreeNode

	if len(iOrder) == 0 || len(pOrder) == 0 {
		return nil
	}

	node = new(TreeNode)
	node.Val = pOrder[len(pOrder) - 1]

	idx := func () int {
		for i, v := range iOrder {
			if v == node.Val {
				return i
			}
		}
		return 0
	}()

	node.Left = buildTree2Order(iOrder[:idx], pOrder[:idx])
	node.Right = buildTree2Order(iOrder[idx + 1:], pOrder[idx:len(pOrder) - 1])

	return node
}

func buildTree2(iOrder []int, pOrder []int) *TreeNode {

	if len(iOrder) != len(pOrder) {
		return nil
	} else if len(iOrder) == 0 {
		return nil
	}

	return buildTree2Order(iOrder, pOrder)
}

func BuildTree2() {
	pOrder := []int{1, 3, 2, 6, 5, 4}
	iOrder := []int{1, 2, 3, 4, 5, 6}

	fmt.Printf("<106> ")
	fmt.Printf("%v\n", buildTree2(iOrder, pOrder))
}

