package problems

import "fmt"
import "math"

func maxProfit(prices []int) int {
	p := prices

	if len(p) == 0 {
		return 0
	}

	min, max := math.MaxInt32, 0

	for _, v := range p {
		if v < min {
			min = v
		}

		if v - min > max {
			max = v - min
		}
	}

	return max
}

func MaxProfit() {
	data := []int {7, 1, 5, 3, 6, 4}

	fmt.Printf("<121> ")
	fmt.Println(maxProfit(data))
}

