package problems

import "fmt"

func generate(numRows int) [][]int {
	n := numRows

	if n == 0 {
		return [][]int{}
	}

	out := make([][]int, n)
	out[0] = []int{1}

	for i := 1; i < n; i++ {
		last := out[i - 1]
		cur := make([]int, i + 1)
		cur[0] = 1

		for k := 1; k < i; k++ {
			cur[k] = last[k - 1] + last[k]
		}

		cur[i] = 1

		out[i] = cur
	}

	return out
}

func Generate() {
	fmt.Printf("<118> ")
	fmt.Println(generate(5))
}

