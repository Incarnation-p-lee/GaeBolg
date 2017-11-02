package problems

import "fmt"

var prev *TreeNode

func flattenOrder(n *TreeNode) {
	if n == nil {
		return
	}

	l, r := n.Left, n.Right

	if prev != nil {
		prev.Right = n
		prev.Left = nil
	}

	prev = n

	flattenOrder(l)
	flattenOrder(r)
}

func flatten(root *TreeNode)  {
	prev = nil

	flattenOrder(root)
}

func Flatten()  {
	var t1, t2, t3, t4, t5, t6 TreeNode

	t1.Val, t2.Val, t3.Val = 1, 2, 3
	t2.Left, t2.Right = &t1, &t3
	t1.Left, t1.Right, t3.Left, t3.Right = nil, nil, nil, nil

	t4.Val, t5.Val, t6.Val = 5, 4, 6
	t5.Left, t5.Right = &t4, &t6
	t4.Left, t4.Right, t6.Left, t6.Right = nil, nil, nil, nil

	fmt.Printf("<114> ") 
	flatten(&t2)
	fmt.Println(t2)
}

