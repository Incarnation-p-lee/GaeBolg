package problems

import "fmt"

func maxProfit4(k int, prices []int) int {
	p := prices

	if k == 0 || len(p) == 0 {
		return 0
	}

	n := len(p)

	if k > n / 2 {
		return maxProfit2(prices)
	}

	dpLast, dpTotal := make([][]int, n), make([][]int, n)

	for i, _ := range dpLast {
		dpLast[i] = make([]int, k + 1)
		dpTotal[i] = make([]int, k + 1)
	}

	var maxInt func (a, b int) int
	maxInt = func (a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	/*
	 * Last[i - 1][j] and diff can share one buy and sell at j.
	 * Last[i][j] = max(Last[i - 1][j] + diff, Total[i - 1][j - 1] + max(0, diff)
	 * Total[i][j] = max(Total[i - 1][j], Last[i][j])
	 */

	for i := 1; i < n; i++ {
		diff := p[i] - p[i - 1]

		for j := 1; j < k + 1; j++ {
			dpLast[i][j] = maxInt(dpLast[i - 1][j] + diff,
				dpTotal[i - 1][j - 1] + maxInt(0, diff))
			dpTotal[i][j] = maxInt(dpTotal[i - 1][j], dpLast[i][j])
		}
	}

	return dpTotal[n - 1][k]
}

func MaxProfit4() {
	data := []int {1, 4, 2, 6, 3, 7, 1, 9}

	fmt.Printf("<188> ")
	fmt.Println(maxProfit4(3, data))
}

