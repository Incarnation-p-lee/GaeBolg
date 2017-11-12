package problems

import "fmt"

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	} else if m == 1 || n == 1 {
		return 1
	}

	min := n

	if m < n {
		min = m
	}

	max := m + n - 2
	min = min - 1

	k, limit := max - min, min + 1

	if k > min {
		k = min
		limit = max - min + 1
	}

	a, b := int64(1), int64(1)

	for limit <= max {
		a *= int64(limit)
		limit++
	}

	for k > 0 {
		b *= int64(k)
		k--
	}

	return int(a / b)
}

func UniquePaths() {
	fmt.Printf("<062> ")
	fmt.Println(uniquePaths(50, 2))
}


