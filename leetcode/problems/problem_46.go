package problems

import "fmt"

func permuteDfs(nums []int, deep int, buf *[]int, out *[][]int, visit []bool) {
	if deep == len(nums) {
		data := make([]int, len(*buf))
		copy(data, *buf)
		*out = append(*out, data)
	} else {
		for i, v := range nums {
			if !visit[i] {
				*buf = append(*buf, v)
				visit[i] = true

				permuteDfs(nums, deep + 1, buf, out, visit)

				*buf = (*buf)[:len(*buf) - 1]
				visit[i] = false
			}
		} 
	}
}

func permute(nums []int) [][]int {
	var buf []int
	var out [][]int
	var visit []bool = make([]bool, len(nums))

	permuteDfs(nums, 0, &buf, &out, visit)

	fmt.Println(out)

	return out
}

func Permute() {
	nums := []int {1, 2, 3}
 
	permute(nums)
}

