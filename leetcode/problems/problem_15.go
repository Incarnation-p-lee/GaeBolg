package problems

import "fmt"
import "sort"

func threeSum(nums []int) [][]int {
	var out [][]int

	sort.Ints(nums)

	for i, v := range nums {
		a := v

		if i > 0 && v == nums[i - 1] {
			continue
		}

		j, k := i + 1, len(nums) - 1

		for j < k {
			b, c := nums[j], nums[k]
			if j > i + 1 && b == nums[j - 1] {
				j++
				continue
			}
			if k < len(nums) - 1 && c == nums[k + 1] {
				k--
				continue
			}

			if a + b + c < 0 {
				j++
			} else if a + b + c > 0 {
				k--
			} else {
				cad := make([]int, 0)
				cad = append(cad, a, b, c)
				out = append(out, cad)
				j++
				k--
			}
		}
	}

	return out
}

func ThreeSum() {
	var data []int = []int{-2, -2, 0, 0, 2, 2, 2}

	fmt.Printf("<015> ")
	fmt.Println(threeSum(data))
}

