package problems

import "fmt"

func generateMatrix(n int) [][]int {
	if n < 1 {
		return [][]int{}
	}

	out := make([][]int, n)
	for k := 0; k < n; k++ {
		out[k] = make([]int, n)
	}

	i, j, x := 0, 0, 1
	direct := ">"
	l, r, u, d := 0, n - 1, 0, n - 1

	for i >= l && j <= r && i >= u && j <= d && l <= r && u <= d {
		out[i][j] = x

		switch direct {
		case ">":
			if j == r {
				direct = "v"
				i++
				u++
			} else {
				j++
			}
		case "v":
			if i == d {
				direct = "<"
				j--
				r--
			} else {
				i++
			}
		case "<":
			if j == l {
				direct = "^"
				i--
				d--
			} else {
				j--
			}
		case "^":
			if i == u {
				direct = ">"
				j++
				l++
			} else {
				i--
			}
		}

		x++
	}

	return out
}

func GenerateMatrix() {
	fmt.Printf("<059> ")
	fmt.Println(generateMatrix(4))
}

