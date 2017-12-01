package problems

import "fmt"

func findMinDuplicated(nums []int) int {
	n := nums

	if len(n) == 0 {
		return 0
	}

	l, r := 0, len(n) - 1

	for l < r {
		if n[l] == n[r] {
			v := n[r]

			for r > l {
				if n[r] != v {
					break
				}
				r--
			}

			if l == r {
				return v
			}
		}

		m := (l + r) / 2
		v := n[m]

		if n[l] <= v && v <= n[r] {
			return n[l]
		}

		if v < n[l] {
			r = m
		} else if v > n[r] {
			l = m + 1
		}
	}

	return n[r]
}

func FindMinDuplicated() {
	data := []int {3, 3, 1, 3, 3}

	fmt.Printf("<154> ")
	fmt.Println(findMinDuplicated(data))
}

