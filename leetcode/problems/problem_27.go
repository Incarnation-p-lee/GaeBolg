package problems

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	k, count := 0, 0

	for _, v := range nums {
		if v != val {
			nums[k] = v
			k++
			count++
		}
	}

	return count
}

func RemoveElement() {
	fmt.Printf("<027> ")
	fmt.Println(removeElement([]int{1, 1, 3, 2, 3, 4}, 3))
}

