package problems

import "fmt"
import "math"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return math.MinInt32
	}

	max, sum := math.MinInt32, 0

	for _, v := range nums {
		if sum < 0 {
			sum = 0
		}

		sum += v

		if sum > max {
			max = sum
		}
	}

	return max
}

func MaxSubArray() {
	fmt.Printf("<053> ")
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

