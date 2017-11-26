package problems

import "fmt"

func longestConsecutive(nums []int) int {
	n := nums

	if len(nums) == 0 {
		return 0
	}

	max := 0
	vmap := make(map[int]int)
	visited := make([]bool, len(n))

	for i, v := range n {
		vmap[v] = i
	}

	for i, v := range n {
		if !visited[i] {
			count, d := 0, v

			for {
				j, ok := vmap[d]

				if !ok {
					break
				}

				d++
				count++
				visited[j] = true
			}

			d = v - 1

			for {
				j, ok := vmap[d]

				if !ok {
					break
				}

				d--
				count++
				visited[j] = true
			}

			if count > max {
				max = count
			}
		}
	}

	return max
}

func LongestConsecutive() {
	fmt.Printf("<128> ")
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

