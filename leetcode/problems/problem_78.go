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

func subsets2Dfs(d []int, idx int, buf *[]int, out *[][]int) {
	if (idx == len(d)) {
		m := make([]int, len(*buf))
		copy(m, *buf)
		*out = append(*out, m)
	} else {
		*buf = append(*buf, d[idx])
		subsets2Dfs(d, idx + 1, buf, out)
		*buf = (*buf)[:len(*buf) - 1]

		subsets2Dfs(d, idx + 1, buf, out)
	}
}

func subsets2(nums []int) [][]int {
	var out [][]int
	var buf []int

	subsets2Dfs(nums, 0, &buf, &out)

	fmt.Println(out)

	return out
}

func Subsets() {
	data := []int{1, 2, 3}

	fmt.Printf("<78> ")
	subsets(data)

	fmt.Printf("<78> ")
    subsets2(data)
}

