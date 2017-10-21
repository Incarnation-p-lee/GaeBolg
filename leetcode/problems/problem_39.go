package problems

import "fmt"
import "sort"

/* Problem 39. Combination Sum */
func CombinationSum() {
	candidates := []int{42,26,36,38,35,41,20,47,45,23,33,39,25,43,29,31,28,48,21,46,22,30,37,32,44,40}
	target := 55

	combinationSum(candidates, target)
}

func combinationDfs(in []int, i int, t int, buf []int, k int, out *[][]int) {
	d := in[i]

	if i == len(in) || t < d {
		return
	} else if i > 0 && in[i] == in[i - 1] {
		/* if sorted are same value, first one will cover next */
		return
	} else if t == d {
		buf[k] = d
		m := make([]int, k + 1)

		copy(m, buf) /* copy instead of assignment */
		*out = append(*out, m)
	} else if t > d {
		t = t - d
		buf[k] = d

		for j := i; j < len(in); j++ { /* from i to rest of slice */
			combinationDfs(in, j, t, buf, k + 1, out)
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
	var out [][]int

	buf := make([]int, target)
	sort.Ints(candidates)

	for i, _ := range candidates {
		/* for each element can be the start one */
		combinationDfs(candidates, i, target, buf, 0, &out)
	}

	fmt.Println(out)

	return out
}

