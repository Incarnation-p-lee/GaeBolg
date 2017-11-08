package problems

import "fmt"

func firstMissingPositive(nums []int) int {
	var nmap = make(map[int]int)

	for _, v := range nums {
		if v > 0 {
			nmap[v] = v
		}
	}

	i := 1

	for {
		if _, ok := nmap[i]; !ok {
			break
		}
		i++
	}

	return i
}

func FirstMissingPositive() {
	data := []int{2, 1, 4, 6, 5, 3, 8, 0}

	fmt.Printf("<041> ")
	fmt.Println(firstMissingPositive(data))
}

