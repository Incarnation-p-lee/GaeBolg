package problems

import "fmt"

func combineDfs(d []int, idx int, k int, buf *[]int, out *[][]int) {
	if (len(*buf) == k) {
		m := make([]int, k)
		copy(m, *buf)
		*out = append(*out, m)
	} else {
		for j := idx; j < len(d); j++ {
			v := d[j]

			*buf = append(*buf, v)
			combineDfs(d, j + 1, k, buf, out)
			*buf = (*buf)[:len(*buf) - 1]
		}
	}
}

func combine(n int, k int) [][]int {
	var buf []int
	var out [][]int
	var data []int

	for i := 0; i < n; i++ {
		data = append(data, i + 1)
	}

	combineDfs(data, 0, k, &buf, &out)

	return out
}

func Combine() {
	n := 4
	k := 2

	fmt.Printf("<077> ")
	fmt.Println(combine(n, k))
}

