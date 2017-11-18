package problems

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	d := matrix

	if len(d) == 0 {
		return 0
	} else if len(d[0]) == 0 {
		return 0
	}

	max := 0
	h := make([]int, len(d[0]))

	var maxInt func (a, b int) int
	maxInt = func (a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	var computeHeight func (i int)
	computeHeight = func (i int) {
		for k := 0; k < len(d[0]); k++ {
			for j := i; j < len(d); j++ {
				if j == i {
					h[k] = int(d[i][k] - '0')
				} else if d[j][k] != '0' && h[k] != 0 {
					h[k]++
				} else {
					break
				}
			}
		}
	}

	for i, _ := range d {
		computeHeight(i)
		max = maxInt(largestRectangleArea2(h), max)
	}

	return max
}

func MaximalRectangle() {
	data := [][]byte {
		[]byte {'1', '0', '1', '0', '0',},
		[]byte {'1', '0', '1', '1', '1',},
		[]byte {'1', '1', '1', '1', '1',},
		[]byte {'1', '1', '1', '1', '1',},
	}

	fmt.Printf("<085> ")
	fmt.Println(maximalRectangle(data))
}

