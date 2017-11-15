package problems

import "fmt"

func setZeroes(matrix [][]int) {
	d := matrix

	if len(d) == 0 {
		return
	}

	isRowZero, isColZero := false, false

	for _, v := range d[0] {
		if v == 0 {
			isRowZero = true
			break
		}
	}

	for _, v := range d {
		if v[0] == 0 {
			isColZero = true
			break
		}
	}

	for i := 1; i < len(d); i++ {
		for j := 1; j < len(d[i]); j++ {
			if d[i][j] == 0 {
				d[0][j] = 0
				d[i][0] = 0
			}
		}
	}

	for j := 1; j < len(d[0]); j++ {
		if d[0][j] == 0 {
			for i := 0; i < len(d); i++ {
				d[i][j] = 0
			}
		}
	}

	for i := 1; i < len(d); i++ {
		if d[i][0] == 0 {
			for j := 0; j < len(d[0]); j++ {
				d[i][j] = 0
			}
		}
	}

	if isRowZero {
		for j := 0; j < len(d[0]); j++ {
			d[0][j] = 0
		}
	}

	if isColZero {
		for i := 0; i < len(d); i++ {
			d[i][0] = 0
		}
	}
}

func SetZeroes() {
	data := [][]int {
		[]int {1, 1, 1, 0, 1},
		[]int {1, 1, 1, 2, 1},
		[]int {1, 1, 1, 3, 1},
		[]int {1, 1, 0, 4, 1},
		[]int {1, 1, 1, 0, 1},
		[]int {1, 0, 1, 1, 1},
	}

	fmt.Printf("<073> ")
	setZeroes(data)
	fmt.Println(data)
}

