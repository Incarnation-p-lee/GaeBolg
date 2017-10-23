package problems

import "fmt"
import "strconv"

func getPermutationDfs(n []int, dp int, k *int, buf *[]int, visit []bool) bool {
	if dp == len(n) {
		(*k)--
		if *k == 0 {
			return true
		}
	} else {
		for i, v := range n {
			if !visit[i] {
				*buf = append(*buf, v)
				visit[i] = true

				if ok := getPermutationDfs(n, dp + 1, k, buf, visit); ok {
					return true
				}

				*buf = (*buf)[:len(*buf) - 1]
				visit[i] = false 
			}
		}
	}

	return false
}

func getPermutation(n int, k int) string {
	var s string
	var buf []int
	var nums []int = make([]int, n)
	var visit []bool = make([]bool, n)

	for i, _ := range nums {
		nums[i] = i + 1
	}

	count := k
	getPermutationDfs(nums, 0, &count, &buf, visit)

	for _, v := range buf {
		s += strconv.Itoa(v)
	}

	fmt.Println(s)

	return s
}


func GetPermutation() {
	n := 6
	k := 23

	fmt.Printf("<60> ")
	getPermutation(n, k)
}

