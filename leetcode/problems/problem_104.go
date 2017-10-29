package problems

import "fmt"

func maxDepth(root *TreeNode) int {
	var depth int
	var stack, slave []*TreeNode

	if root == nil {
		return 0
	}

	stack = append(stack, root)

	for {
		if len(stack) == 0 {
			depth++

			if len(slave) == 0 {
				break
			}

			stack = make([]*TreeNode, len(slave))
			copy(stack, slave)
			slave = slave[:0]
		}

		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		if node.Left != nil {
			slave = append(slave, node.Left)
		}

		if node.Right != nil {
			slave = append(slave, node.Right)
		}
	}

	return depth
}

func MaxDepth() {
	var t1, t2, t3, t4, t5, t6, t7 TreeNode

	t1.Val, t2.Val, t3.Val = 2, 4, 6
	t2.Left, t2.Right = &t1, &t3

	t4.Val, t5.Val, t6.Val, t7.Val = 1, 3, 5, 7
	t1.Left, t1.Right = &t4, &t5
	t3.Left, t3.Right = &t6, &t7
	t4.Left, t4.Right, t5.Left, t5.Right = nil, nil, nil, nil
	t6.Left, t6.Right, t7.Left, t7.Right = nil, nil, nil, nil

	fmt.Printf("<103> ")
	fmt.Println(maxDepth(&t2)) 
}

