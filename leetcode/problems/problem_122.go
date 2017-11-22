package problems

import "fmt"

func maxProfit2(prices []int) int {
	p := prices

	if len(prices) == 0 {
		return 0
	}

	max := 0

	for i := 0; i < len(p); i++ {
		end := len(p) - 1

		for k := i + 1; k < len(p); k++ {
			if p[k] < p[k - 1] {
				end = k - 1
				break
			}
		}

		if end > i {
			max += p[end] - p[i]
			i = end
		}
	}

	return max
}

func MaxProfit2() {
	data := []int {7, 1, 5, 3, 6, 4}

	fmt.Printf("<122> ")
	fmt.Println(maxProfit2(data))
}

