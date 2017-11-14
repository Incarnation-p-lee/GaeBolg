package problems

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	data, t := matrix, target

	if len(data) == 0 || len(data[0]) == 0 {
		return false
	}

	i, j := 0, len(data) - 1

	for i <= j {
		m := (i + j) / 2

		if t < data[m][0] {
			j = m - 1
		} else if t > data[m][0] {
			i = m + 1
		} else {
			return true
		}
	}

	k := j
	if j < 0 {
		k = 0
	}

	i, j = 0, len(data[k]) - 1

	for i <= j {
		m := (i + j) / 2

		if t > data[k][m] {
			i = m + 1
		} else if t < data[k][m] {
			j = m - 1
		} else {
			return true
		}
	}

	return false
}

func SearchMatrix() {
	data := [][]int {
		[]int {1,   3,  5,  7},
		[]int {10, 11, 16, 20},
		[]int {23, 30, 34, 50},
	}

	fmt.Printf("<074> ")
	fmt.Println(searchMatrix(data, 33))
}

