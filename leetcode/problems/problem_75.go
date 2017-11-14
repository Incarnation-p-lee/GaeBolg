package problems

import "fmt"

func sortColors(nums []int) {
	d := nums

	if len(d) == 0 {
		return
	}

	m := make([]int, 3)

	for _, v := range nums {
		m[v]++
	}

	k := 0

	for i, limit := range m {
		for j := 0; j < limit; j++ {
			nums[k] = i
			k++
		}
	}
}

func SortColors() {
	data := []int {0, 2, 0, 1, 2, 0, 1, 0, 2, 1, 0, 0, 2, 1, 2, 1}

	fmt.Printf("<075> ")
	sortColors(data)
	fmt.Println(data)
}

