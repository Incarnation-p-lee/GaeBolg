package problems

import "fmt"
import "sort"

func permuteUniqueDfs(nums []int, buf *[]int, out *[][]int, deep int,
	 visit []bool) {
	if deep == len(nums) {
		data := make([]int, len(*buf))
		copy(data, *buf)
		*out = append(*out, data)
	} else {
		for i, v := range nums {
			if !visit[i] && i > 0 && !visit[i - 1] && nums[i - 1] == nums[i] {
				/*
				 * As input sorted, if n[i - 1] == n[i] && i - 1 is not visited,
				 * We can safely skip n[i] Dfs, it will be covered by i - 1.
				 */
				continue
			} else if !visit[i] {
				*buf = append(*buf, v)
				visit[i] = true

				permuteUniqueDfs(nums, buf, out, deep + 1, visit)

				*buf = (*buf)[:len(*buf) - 1]
				visit[i] = false
			}
		} 
	}
}

func permuteUnique(nums []int) [][]int {
	var buf []int
	var out [][]int
	var visit []bool = make([]bool, len(nums))

	sort.Ints(nums)
    
	permuteUniqueDfs(nums, &buf, &out, 0, visit)

	fmt.Println(out)

	return out
}

func PermuteUnique() {
	nums := []int{3, 3, 0, 3, 0}


	permuteUnique(nums)    
}

