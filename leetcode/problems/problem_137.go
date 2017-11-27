package problems

import "fmt"

func singleNumber2(nums []int) int {
	n := nums

	if len(n) == 0 {
		return 0
	}

	one, two, three := 0, 0, 0

	for _, v := range nums {
		tmp := one
		one = one ^ v
		two = two | (tmp & v)
		three = two & one

		one = one & (^three)
		two = two & (^three)
	}

	return one
}

func SingleNumber2() {
	data := []int {1, 1, 1, 2, 2, 2, 3, 4, 4, 4}

	fmt.Printf("<137> ")
	fmt.Println(singleNumber2(data))
}

