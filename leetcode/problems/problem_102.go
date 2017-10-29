package problems

import "fmt"

func levelOrder(root *TreeNode) [][]int {
	var buf []int
	var out [][]int    
	var queue, slave []*TreeNode

	if root == nil {
		return out
	}

	queue = append(queue, root)

	for {
		if len(queue) == 0 {
			m := make([]int, len(buf))
			copy(m, buf)
			out = append(out, m)

			if len(slave) == 0 {
				break
			}

			queue = make([]*TreeNode, len(slave))
			copy(queue, slave)
			buf = buf[:0]
			slave = slave[:0]
		}

		node := queue[0]
		queue = queue[1:len(queue)]
		buf = append(buf, node.Val)

		if node.Left != nil {
			slave = append(slave, node.Left)
		}

		if node.Right != nil {
			slave = append(slave, node.Right)
		}
	}

	return out
}

func LevelOrder() {
	var t1, t2, t3, t4, t5, t6, t7 TreeNode

	t1.Val, t2.Val, t3.Val = 2, 4, 6
	t2.Left, t2.Right = &t1, &t3

	t4.Val, t5.Val, t6.Val, t7.Val = 1, 3, 5, 7
	t1.Left, t1.Right = &t4, &t5
	t3.Left, t3.Right = &t6, &t7
	t4.Left, t4.Right, t5.Left, t5.Right = nil, nil, nil, nil
	t6.Left, t6.Right, t7.Left, t7.Right = nil, nil, nil, nil

	fmt.Printf("<102> ")
	fmt.Println(levelOrder(&t2))
}

