package problems

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	g, c := gas, cost

	if len(g) != len(c) || len(g) == 0 {
		return -1
	}

	diff := make([]int, len(g))

	for i, _ := range diff {
		diff[i] = g[i] - c[i]
	}

	total, sum, idx := 0, 0, 0

	for i, v := range diff {
		total += v

		if sum < 0 {
			sum = v
			idx = i
		} else {
			sum += v
		}
	}

	if total < 0 {
		return -1
	} else {
		return idx
	}
}

func CanCompleteCircuit() {
	gas := []int{2, 4, 4, 5, 6, 3}
	cost := []int{3, 1, 2, 5, 8, 1}

	fmt.Printf("<134> ")
	fmt.Println(canCompleteCircuit(gas, cost))
}

