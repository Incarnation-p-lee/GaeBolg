package problems

import "fmt"

func getRow(rowIndex int) []int {
	n := rowIndex
	d := make([]int, n + 1)
	d[0] = 1

	for k := 1; k < n + 1; k++ {
		l := k + 1
		last := d[0]

		for i := 1; i < l - 1; i++ {
			tmp := d[i]
			d[i] = last + d[i]
			last = tmp
		}

		d[l - 1] = 1
	}

	return d
}

func GetRow() {
	fmt.Printf("<119> ")
	fmt.Println(getRow(3))
}

