package problems

import "fmt"
import "sort"

func fourSum(nums []int, target int) [][]int {
	var out [][]int

	sort.Ints(nums)

	for i, v := range nums {
		a := v

		if i > 0 && v == nums[i - 1] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			b := nums[j]

			if j > i + 1 && b == nums[j - 1] {
				continue
			}

			k, l := j + 1, len(nums) - 1

			for k < l {
				c, d := nums[k], nums[l]
				m := a + b + c + d

				if k > j + 1 && c == nums[k - 1] {
					k++
					continue
				}

				if l < len(nums) - 1 && d == nums[l + 1] {
					l--
					continue
				}

				if m > target {
					l--
				} else if m < target {
					k++
				} else {
					cad := make([]int, 0)
					cad = append(cad, a, b, c, d)
					out = append(out, cad)
					l--
					k++
				}
			}
		}
	}

	return out
}

func FourSum() {
	var nums []int
	var target int

	nums = []int{1, 0, -1, 0, -2, 2}
	target = 0

	fmt.Printf("<018> ")
	fmt.Println(fourSum(nums, target))
}

