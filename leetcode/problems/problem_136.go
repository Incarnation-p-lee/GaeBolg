package problems

import "fmt"

func singleNumber(nums []int) int {
	n := nums

	if len(n) == 0 {
		return 0
	}

	single := 0

	for _, v := range n {
		single = single ^ v
	}

	return single
}

func SingleNumber() {
	data := []int {1, 2, 3, 3, 2, 1, 9}

	fmt.Printf("<136> ")
	fmt.Println(singleNumber(data))
}

