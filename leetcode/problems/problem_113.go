package problems

import "fmt"

func pathSumDfs(root *TreeNode, sum int, buf *[]int, out *[][]int) {
	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			m := make([]int, len(*buf))
			copy(m, *buf)
			m = append(m, root.Val)
			*out = append(*out, m)
		}
	} else {
		if root.Left != nil {
			*buf = append(*buf, root.Val)
			pathSumDfs(root.Left, sum - root.Val, buf, out)
			*buf = (*buf)[:len(*buf) - 1]
		}

		if root.Right != nil {
			*buf = append(*buf, root.Val)
			pathSumDfs(root.Right, sum - root.Val, buf, out)
			*buf = (*buf)[:len(*buf) - 1]
		}
	}
}

func pathSum(root *TreeNode, sum int) [][]int {
	var out [][]int
	var buf []int

	if root == nil {
		return out
	}

	pathSumDfs(root, sum, &buf, &out)

	return out
}

func PathSum() {
	var t1, t2, t3, t4, t5, t6, t7 TreeNode

	t1.Val, t2.Val, t3.Val = 2, 4, 6
	t2.Left, t2.Right = &t1, &t3

	t4.Val, t5.Val, t6.Val, t7.Val = 1, 3, 5, 7
	t1.Left, t1.Right = &t4, &t5
	t3.Left, t3.Right = &t6, &t7
	t4.Left, t4.Right, t5.Left, t5.Right = nil, nil, nil, nil
	t6.Left, t6.Right, t7.Left, t7.Right = nil, nil, nil, nil

	fmt.Printf("<113> ")
	fmt.Println(pathSum(&t2, 15))
}

