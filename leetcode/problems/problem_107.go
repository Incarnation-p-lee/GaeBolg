package problems

import "fmt"

func levelOrderBottom(root *TreeNode) [][]int {
	var buf []int
	var out, tmp [][]int 
	var queue, slave []*TreeNode

	if root == nil {
		return out 
	}

	queue = append(queue, root)

	for {
		if len(queue) == 0 {
			data := make([]int, len(buf))
			copy(data, buf)
			tmp = append(tmp, data)
			buf = buf[:0]

			if len(slave) == 0 {
				break
			}

			queue = make([]*TreeNode, len(slave))
			copy(queue, slave)
			slave = slave[:0]
		}

		node := queue[0]
		queue = queue[1:]
		buf = append(buf, node.Val) 

		if node.Left != nil {
			slave = append(slave, node.Left)
		}

		if node.Right != nil {
			slave = append(slave, node.Right)
		}
	}

	for i := len(tmp) - 1; i >= 0; i-- {
		out = append(out, tmp[i])
	}

	return out
}

func LevelOrderBottom() {
	var t1, t2, t3, t4, t5, t6, t7 TreeNode

	t1.Val, t2.Val, t3.Val = 2, 4, 6
	t2.Left, t2.Right = &t1, &t3

	t4.Val, t5.Val, t6.Val, t7.Val = 1, 3, 5, 7
	t1.Left, t1.Right = &t4, &t5
	t3.Left, t3.Right = &t6, &t7
	t4.Left, t4.Right, t5.Left, t5.Right = nil, nil, nil, nil
	t6.Left, t6.Right, t7.Left, t7.Right = nil, nil, nil, nil

	fmt.Printf("<107> ")
	fmt.Println(levelOrderBottom(&t2))
}

