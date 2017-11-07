package problems

import "fmt"
import "sort"

func nextPermutation(nums []int)  {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}

	k := 0

	for i := len(nums) - 1; i > 0; i-- {
		if nums[i - 1] >= nums[i] {
			k = i
			break
		}
	}

	if k != len(nums) - 1 {
		l := len(nums)
		nums[l - 2], nums[l - 1] = nums[l - 1], nums[l - 2]
	} else {
		for i := k - 1; i > 0; i-- {
			if nums[i - 1] < nums[i] {
				m := k - 1
				for p := k - 1; p >= i; p-- {
					if nums[p] > nums[i - 1] {
						m = p
						break
					}
				}

				if nums[k] < nums[m] && nums[k] > nums[i - 1] {
					m = k
				}

				nums[i - 1], nums[m] = nums[m], nums[i - 1]
				sort.Ints(nums[i:])
				return
			}
		}

		sort.Ints(nums)
	}
}

func NextPermutation() {
	nums := []int{2,3,0,2,4,1}

	fmt.Printf("<031> ")
	nextPermutation(nums)

	fmt.Println(nums)
}

