package problems

import "fmt"

func twoSum(nums []int, target int) []int {
	var out []int
	var imap = make(map[int]int)

	for i, v := range nums {
		imap[v] = i
	}

	for i, v := range nums {
		if j, ok := imap[target - v]; ok && i != j {
			return append(out, i, j)
		}
	}

	return out;
}

func TwoSum() {
	var n []int = []int{-1, -2, -3, -4, -5}
	var t int = -8

	fmt.Printf("<001> ")
	fmt.Println(twoSum(n, t))
}

