package problems

import "fmt"

func findMin(nums []int) int {
	n := nums

	if len(n) == 0 {
		return 0
	}

	l, r := 0, len(n) - 1

	for l < r {
		if n[l] < n[r] {
			return n[l]
		}

		m := (l + r) / 2
		v := n[m]

		if n[l] > v {
			r = m
		} else if v > n[r] {
			l = m + 1
		}
	}

	return n[r]
}

func FindMin() {
	data := []int {4, 5, 6, 7, 0, 1, 2}

	fmt.Printf("<153> ")
	fmt.Println(findMin(data))
}
