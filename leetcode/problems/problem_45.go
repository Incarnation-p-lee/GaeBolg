package problems

import "fmt"
import "math"

func jump(nums []int) int {
	steps := make([]int, len(nums))

	steps[0] = 0

	for i := 1; i < len(nums); i++ {
		min := math.MaxInt32

		for j := 0; j < i; j++ {
			d := j + nums[j]

			if d >= i && steps[j] + 1 < min {
				min = steps[j] + 1
			}
		}

		steps[i] = min
	}

	return steps[len(nums) - 1]
}

func jump2(nums []int) int {
	var last, ret, cur int

	last, ret, cur = 0, 0, 0

	for i, v := range nums {
		if i > last {
			ret++
			last = cur
		}

		if i + v > cur {
			cur = i + v
		}
	}

	return ret
}

func Jump() {
	fmt.Printf("<045> ")
	fmt.Println(jump([]int{2, 3, 1, 1, 4, 3, 2, 4, 1}))

	fmt.Printf("<045> ")
	fmt.Println(jump2([]int{2, 3, 1, 1, 4, 3, 2, 4, 1}))
}

