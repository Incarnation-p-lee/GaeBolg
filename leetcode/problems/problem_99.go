package problems

import "fmt"

var lastNode *TreeNode

func recoverTreeInOrder(root *TreeNode, f, s, n **TreeNode) {
	if root == nil {
		return
	}

	recoverTreeInOrder(root.Left, f, s, n)

	if root != lastNode && root.Val < lastNode.Val && *f == nil {
		*f = lastNode
		*n = root
	} else if root != lastNode && *f != nil && root.Val < lastNode.Val {
		*s = root
	}

	lastNode = root

	recoverTreeInOrder(root.Right, f, s, n)
}

func recoverTree(root *TreeNode)  {
	var first, next, second *TreeNode

	if root == nil {
		return
	}

	lastNode = root

	for ; lastNode.Left != nil; {
		lastNode = lastNode.Left;
	}

	recoverTreeInOrder(root, &first, &second, &next)

	if first == nil {
		return
	}

	if second != nil {
		first.Val, second.Val = second.Val, first.Val
	} else {
		first.Val, next.Val = next.Val, first.Val
	}
}


func RecoverTree()  {
	var t1, t2, t3, t4, t5, t6, t7 TreeNode

	t1.Val, t2.Val, t3.Val, t4.Val, t5.Val, t6.Val, t7.Val = 1, 2, 3, 4, 5, 6, 7

	t4.Left, t4.Right = &t2, &t6
	t2.Left, t2.Right = &t1, &t3
	t6.Left, t6.Right = &t5, &t7

	t1.Left, t1.Right = nil, nil
	t3.Left, t3.Right = nil, nil
	t5.Left, t5.Right = nil, nil
	t7.Left, t7.Right = nil, nil

	t4.Val, t3.Val = t3.Val, t4.Val

	fmt.Printf("<099> ")
	fmt.Printf("%v -> ", t4)
	recoverTree(&t4)
	fmt.Printf("%v\n", t4)
}

