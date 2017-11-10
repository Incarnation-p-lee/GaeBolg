package problems

import "fmt"

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	last, current := 0, 0

	for i, v := range nums {
		if i > last {
			if current <= last {
				return false
			}

			last = current
		}

		if i + v > current {
			current = i + v
		}
	}

	return true
}

func CanJump() {
	data := []int{3, 2, 1, 0, 4}

	fmt.Printf("<055> ")
	fmt.Println(canJump(data))
}

