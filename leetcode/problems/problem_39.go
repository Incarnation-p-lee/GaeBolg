package problems

import "fmt"
import "sort"

/* Problem 39. Combination Sum */
func CombinationSum() {
	candidates := []int{2, 3, 6, 7}
	target := 7

	combinationSum(candidates, target)
}

func combinationDfs(in []int, i int, t int, buf *[]int, out *[][]int) {

	if t < 0 {
		return
	} else if t == 0 {
        m := make([]int, len(*buf))
		copy(m, *buf)                     /* copy instead of assignment */
		*out = append(*out, m)
	} else {
		for j := i; j < len(in); j++ {    /* from i to rest of slice */
			d := in[j]
			*buf = append(*buf, d)        /* push d */

			combinationDfs(in, j, t - d, buf, out)

			*buf = (*buf)[:len(*buf) - 1] /* pop d */
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
	var buf []int
	var out [][]int

	sort.Ints(candidates)

	combinationDfs(candidates, 0, target, &buf, &out)

	fmt.Println(out)

	return out
}

