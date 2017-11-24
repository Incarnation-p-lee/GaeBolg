package problems

import "fmt"
import "math"

func maxProfit3(prices []int) int {
	p := prices

	if len(p) == 0 {
		return 0
	}

	dpLeft, dpRight := make([]int, len(p)), make([]int, len(p))
	min, max, pmax := math.MaxInt32, 0, 0

	for i, v := range p {
		if v < min {
			min = v
		}

		val := v - min

		if val > pmax {
			pmax = val
		}

		dpLeft[i] = pmax
	}

	pmax = 0

	for i := len(p) - 1; i >= 0; i-- {
		v := p[i]
		if v > max {
			max = v
		}

		val := max - v

		if val > pmax {
			pmax = val
		}

		dpRight[i] = pmax
	}

	max = 0

	for i := 0; i < len(p); i++ {
		cad := 0
		if i == len(p) - 1 {
			cad = dpLeft[i]
		} else {
			cad = dpLeft[i] + dpRight[i + 1]
		}

		if cad > max {
			max = cad
		}
	}

	return max
}

func MaxProfit3() {
	fmt.Printf("<125> ")
	fmt.Println(maxProfit3([]int{1, 5, 1, 9, 3}))
}

