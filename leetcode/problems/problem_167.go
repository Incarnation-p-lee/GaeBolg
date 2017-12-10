package problems

import "fmt"

func twoSumSorted(numbers []int, target int) []int {
	n := numbers
	t := target

	if len(n) < 2 {
		return []int {-1, -1}
	}

	l, r := 0, len(n) - 1

	for l < r {
		if n[l] + n[r] > t {
			r--
		} else if n[l] + n[r] < t {
			l++
		} else {
			return []int{l + 1, r + 1}
		}
	}

	return []int {-1, -1}
}

func TwoSumSorted() {
	fmt.Printf("<167> ")
	fmt.Println(twoSumSorted([]int{1, 3, 4, 7, 10}, 10))
}

