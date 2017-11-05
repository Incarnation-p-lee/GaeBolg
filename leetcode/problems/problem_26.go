package problems

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	len, i := 0, 0

	for k, v := range nums {
		if k == 0 || (k > 0 && v != nums[k - 1]) {
			nums[i] = v
			len++
			i++
		}
	}

	return len
}

func RemoveDuplicates() {
	nums := []int {1, 1, 1, 2, 3, 4, 4, 4, 5}

	fmt.Printf("<026> ")
	fmt.Println(removeDuplicates(nums))
}

