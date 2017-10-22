package problems

import "fmt"

func permuteDfs(nums []int, deep int, buf *[]int, out *[][]int, m *map[int]int) {
	if deep == len(nums) {
		data := make([]int, len(*buf))
		copy(data, *buf)
		*out = append(*out, data)
		return
	}

	nMap := *m

	for _, v := range nums {
		if _, ok := nMap[v]; !ok {
			*buf = append(*buf, v)
			nMap[v] = v

			permuteDfs(nums, deep + 1, buf, out, m)

			*buf = (*buf)[:len(*buf) - 1]
			delete(nMap, v)
		}
	} 
}

func permute(nums []int) [][]int {
	var nMap map[int]int = make(map[int]int)
	var buf []int
	var out [][]int

	permuteDfs(nums, 0, &buf, &out, &nMap)

	fmt.Println(out)

	return out
}

func Permute() {
	nums := []int {1, 2, 3}
 
	permute(nums)
}

