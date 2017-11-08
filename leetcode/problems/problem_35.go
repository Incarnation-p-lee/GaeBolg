package problems

import "fmt"

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 || target <= nums[0] {
		return 0
	} else if target > nums[len(nums) - 1] {
		return len(nums)
	}

	for i, v := range nums {
		if i > 0 && v >= target {
			return i
		}
	}

	return 0
}

func SearchInsert() {
	data := []int{1, 3, 5, 7, 8, 9}

	fmt.Printf("<035> ")
	fmt.Println(searchInsert(data, 9))
}

