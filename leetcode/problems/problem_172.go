package problems

import "fmt"

func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}

	out := 0

	for n != 0 {
		out = out + n / 5
		n = n / 5
	}

	return out
}

func TrailingZeroes() {
	fmt.Printf("<172> ")
	fmt.Println(trailingZeroes(13))
}

