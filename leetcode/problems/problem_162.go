package problems

import "fmt"

func findPeakElement(nums []int) int {
	n := nums

	if len(n) == 0 {
		return -1
	} else if len(n) == 1 {
		return 0
	}

	l, r := 0, len(n) - 1

	for l < r {
		m := (l + r) / 2

		if n[m] > n[m + 1] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func FindPeakElement() {
	data := []int {1, 2, 3, 1}

	fmt.Printf("<162> ")
	fmt.Println(findPeakElement(data))
}

