package problems

import "fmt"
import "sort"
import "math"

var min, out int

func threeSumClosest(nums []int, target int) int {
	min = math.MaxInt32

	sort.Ints(nums)

	for i, v := range nums {
		a := v

		j, k := i + 1, len(nums) - 1

		for j < k {
			b, c := nums[j], nums[k]

			d := func () int {
				d := a + b + c - target
				if d < 0 {
					d = -d
				}
				return d
			} ()

			if d < min {
				min = d
				out = a + b + c
			}

			if a + b + c < target {
				j++
			} else if a + b + c > target {
				k--
			} else {
				return target
			}
		}
	}

	return out
}

func ThreeSumClosest() {
	var target int
	var nums []int

	target = 1
	nums = []int{-1, 2, 1, -4}

	fmt.Printf("<016> ")
	fmt.Println(threeSumClosest(nums, target))
}

