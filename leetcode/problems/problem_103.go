package problems

import "fmt"

func zigzagLevelOrder(root *TreeNode) [][]int {
	var buf []int
	var isZigZag bool
	var out [][]int    
	var node *TreeNode
	var data, slave []*TreeNode

	if root == nil {
		return out
	}

	isZigZag = true
	data = append(data, root)

	for {
		if len(data) == 0 {
			m := make([]int, len(buf))
			copy(m, buf)
			out = append(out, m)

			if len(slave) == 0 {
				break
			}

			data = make([]*TreeNode, len(slave))
			copy(data, slave)
			buf = buf[:0]
			slave = slave[:0]
			isZigZag = !isZigZag
		}

		node = data[len(data) - 1]
		data = data[:len(data) - 1]
		buf = append(buf, node.Val)

		if node.Left != nil && isZigZag {
			slave = append(slave, node.Left)
		}

		if node.Right != nil && isZigZag {
			slave = append(slave, node.Right)
		}

		if node.Right != nil && !isZigZag {
			slave = append(slave, node.Right)
		}

		if node.Left != nil && !isZigZag {
			slave = append(slave, node.Left)
		}
	}

	return out
}

func ZigzagLevelOrder() {
	var t1, t2, t3, t4, t5, t6, t7 TreeNode

	t1.Val, t2.Val, t3.Val = 2, 4, 6
	t2.Left, t2.Right = &t1, &t3

	t4.Val, t5.Val, t6.Val, t7.Val = 1, 3, 5, 7
	t1.Left, t1.Right = &t4, &t5
	t3.Left, t3.Right = &t6, &t7
	t4.Left, t4.Right, t5.Left, t5.Right = nil, nil, nil, nil
	t6.Left, t6.Right, t7.Left, t7.Right = nil, nil, nil, nil

	fmt.Printf("<103> ")
	fmt.Println(zigzagLevelOrder(&t2))
}

