package problems

import "fmt"

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	var copyTree func (node *TreeNode) *TreeNode
	copyTree = func (node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		t := new(TreeNode)

		t.Val = node.Val
		t.Left = copyTree(node.Left)
		t.Right = copyTree(node.Right)

		return t
	}

	var val int
	var valInit func ()
	valInit = func () {
		val = 1
	}

	var sortTree func (node *TreeNode)
	sortTree = func (node *TreeNode) {
		if node != nil {
			sortTree(node.Left)
			node.Val = val
			val++
			sortTree(node.Right)
		}
	}

	dp := make([][]*TreeNode, n + 1)
	for i, _ := range dp {
		dp[i] = make([]*TreeNode, 0)
	}

	node := TreeNode{1, nil, nil}

	dp[0] = append(dp[0], nil)
	dp[1] = append(dp[1], &node)

	for i := 2; i < n + 1; i++ {
		for j := 0; j < i; j++ {
			for _, l := range dp[j] {
				for _, r := range dp[i - 1 - j] {
					t := new(TreeNode)
					t.Left = copyTree(l)
					t.Right = copyTree(r)
					dp[i] = append(dp[i], t)
				}
			}
		}
	}

	for _, v := range dp[n] {
		valInit()
		sortTree(v)
	}

	return dp[n]
}

func GenerateTrees() {
	fmt.Printf("<095> ")
	fmt.Println(generateTrees(4))
}

