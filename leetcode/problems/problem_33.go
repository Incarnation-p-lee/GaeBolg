package problems

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	i, j := 0, len(nums) - 1

	var bFind func (d []int, s, e, t int) int
	bFind = func (d []int, s, e, t int) int {
		for s <= e {
			k := (s + e) / 2
			m := d[k]

			if t < m {
				e = k - 1
			} else if t > m {
				s = k + 1
			} else {
				return k
			}
		}
		return -1
	}

	for i <= j {
		k := (i + j) / 2
		med := nums[k]

		if nums[i] <= med { /* sorted */
			if target < med {
				if target > nums[i] {
					return bFind(nums, i, k, target)
				} else if target < nums[i] {
					i = k + 1
					continue
				} else {
					return i
				}
			} else if target > med {
				i = k + 1
				continue
			} else {
				return k
			}
		}

		if med <= nums[j] { /* sorted */
			if target > med {
				if target < nums[j] {
					return bFind(nums, k, j, target)
				} else if target > nums[j] {
					j = k - 1
					continue
				} else {
					return j
				}
			} else if target < med {
				j = k - 1
				continue
			} else {
				return k
			}
		}
	}

	return -1
}

func Search() {
	fmt.Printf("<033> ")
	fmt.Println(search([]int{5, 1, 3}, 3))
}

