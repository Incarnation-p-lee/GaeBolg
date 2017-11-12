package problems

import "fmt"

func mySqrt(x int) int {
	if x <= 0 {
		return 0
	} else if x == 1 {
		return 1
	}

	l, r := 1, x

	for l <= r {
		m := (l + r) / 2
		v := m * m

		if v > x {
			r = m - 1
		} else if v < x {
			l = m + 1
		} else {
			return m
		}
	}

	return r
}

func MySqrt() {
	fmt.Printf("<069> ")
	fmt.Println(mySqrt(201012031))
}

