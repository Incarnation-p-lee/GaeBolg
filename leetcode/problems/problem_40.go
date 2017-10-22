package problems

import "fmt"
import "sort"

func CombinationSum2() {
	candidates := []int{1, 1, 1, 2, 5, 7, 6, 6, 6, 10} 
	target := 12

	combinationSum2(candidates, target)
}

func combinationDfs2(input []int, idx int, t int, buf *[]int, out *[][]int) {
	if t == 0 {
		found := make([]int, len(*buf))
		copy(found, *buf)
		*out = append(*out, found)
	} else {
		for j := idx; j < len(input); j++ {
			if j > idx && input[j] == input[j - 1] {
				continue
			}
			
			d := input[j]
			
			if (t < d) { /* input are sorted, if d > t, rest element > t too */
				break
			}
			
			(*buf) = append((*buf), d) /* push d */
			combinationDfs2(input, j + 1, t - d, buf, out)
			(*buf) = (*buf)[:len(*buf) - 1] /* pod */
		}
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	var buf []int
	var out [][]int

	sort.Ints(candidates)
	combinationDfs2(candidates, 0, target, &buf, &out)

	fmt.Println(out)

	return out;
}


