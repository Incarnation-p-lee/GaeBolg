package problems

import "fmt"

func majorityElement(nums []int) int {
	n := nums

	if len(n) == 0 {
		return 0
	} else if len(n) == 1 {
		return n[0]
	}

	nmap := make(map[int]int)

	for _, v := range n {
		val, ok := nmap[v]

		if ok {
			nmap[v] = val + 1

			if nmap[v] > len(n) / 2 {
				return v
			}
		} else {
			nmap[v] = 1
		}
	}

	return 0
}

func MajorityElement() {
	fmt.Printf("<169> ")
	fmt.Println(majorityElement([]int {1, 2, 1, 2, 1, 3, 1, 1, 4, 1}))
}

