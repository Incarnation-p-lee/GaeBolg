package problems

import "fmt"
import "sort"

func subsets(nums []int) [][]int {
	var t []int
	var out [][]int

	out = append(out, t)

	for _, n := range nums {
		for _, v := range out {
			d := make([]int, len(v))
			copy(d, v)
			d = append(d, n)
			sort.Ints(d)
			out = append(out, d)
		}
	}

	fmt.Println(out)

	return out
}

func Subsets() {
	data := []int{9, 0, 3, 5}

	fmt.Printf("<78> ")
	subsets(data)
}

