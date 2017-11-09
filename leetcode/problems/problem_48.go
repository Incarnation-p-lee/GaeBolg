package problems

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)

	if n < 2 {
		return
	}

	d := matrix

	for i := 0; i < n; i++ {
		for j := i; j < n - i - 1; j++ {
			d[i][j], d[j][n - 1 -i], d[n - 1 -i][n - 1 - j], d[n - 1 -j][i] =
			d[n - 1 -j][i], d[i][j], d[j][n - 1 -i], d[n - 1 -i][n - 1 - j]
			/*
			 * (i, j) -> (j, n - i - 1) -> (n - i - 1, n - j - 1) ->
			 * (n - j - 1, i)
			 */
		}
	}
}

func Rotate() {
	data := [][]int {
		[]int{5,1,9,11},
		[]int{2,4,8,10},
		[]int{13,3,6,7},
		[]int{15,14,12,16},
	}

	fmt.Printf("<048> ")
	rotate(data)
	fmt.Println(data)
}

