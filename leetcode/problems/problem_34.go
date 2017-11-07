package problems

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) < 1 || target > nums[len(nums) - 1] || target < nums[0] {
		return []int{-1, -1}
	}

	left, right := -1, -1

	nums = append(nums, nums[len(nums) - 1] + 1)

	for i, v := range nums {
		if v == target && left == -1 {
			left, right = i, i
		} else if v > target && i > 0 && left != -1 {
			right = i - 1
			break
		}
	}

	return []int{left, right}
}

func SearchRange() {
	var data = []int{5, 7, 7, 8, 8, 8, 8, 8, 10}

	fmt.Printf("<034> ")
	fmt.Println(searchRange(data, 8))
}

