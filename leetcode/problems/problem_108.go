package problems

import "fmt"

func sortedArrayToBSTOrder(nums []int) *TreeNode {
	var node *TreeNode

	if len(nums) == 0 {
		return nil
	}

	node = new(TreeNode)
	idx := len(nums) / 2
	node.Val = nums[idx]

	node.Left = sortedArrayToBSTOrder(nums[:idx])
	node.Right = sortedArrayToBSTOrder(nums[idx + 1:])

	return node
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	} 

	return sortedArrayToBSTOrder(nums)
}

func SortedArrayToBST() {
	data := []int {1, 2, 3, 4, 5, 6, 9}

	fmt.Printf("<108> ")
	fmt.Printf("%v\n", sortedArrayToBST(data))
}

