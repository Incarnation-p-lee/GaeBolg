package problems

import "fmt"

func minPathSum(grid [][]int) int {
	d := grid

	if len(d) == 0 {
		return 0
	}

	dp := make([]int, len(d[0]))

	for i, v := range d[0] {
		if i == 0 {
			dp[i] = v
		} else {
			dp[i] = dp[i - 1] + v
		}
	}

	for i := 1; i < len(d); i++ {
		for j, v := range d[i] {
			if j == 0 {
				dp[j] += v
			} else {
				if dp[j - 1] > dp[j] {
					dp[j] += v
				} else {
					dp[j] = dp[j - 1] + v
				}
			}
		}
	}

	return dp[len(d[0]) - 1]
}

func MinPathSum() {
	data := [][]int {
		[]int {1, 3, 1},
		[]int {1, 5, 1},
		[]int {4, 2, 1},
	}

	fmt.Printf("<064> ")
	fmt.Println(minPathSum(data))
}

