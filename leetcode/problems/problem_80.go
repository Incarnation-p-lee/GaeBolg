package problems

import "fmt"

func removeDuplicates2(nums []int) int {
	d := nums

	if len(d) == 0 {
		return 0
	}

	k, count := 0, 1

	for i, v := range d {
		if i != 0 {
			if v != d[k - 1] {
				count = 1
			} else if count > 1 {
				continue
			} else {
				count++
			}
		}

		d[k] = v
		k++
	}

	fmt.Println(d)

	return k
}

func RemoveDuplicates2() {
	data := []int{1, 1, 1, 2, 2, 3, 4, 4, 4}

	fmt.Printf("<080> ")
	fmt.Println(removeDuplicates2(data))
}

