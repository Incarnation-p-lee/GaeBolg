package problems

import "fmt"

func minimumTotal(triangle [][]int) int {
	data := triangle

	if len(data) == 0 {
		return 0
	}

	var minInt func (a, b int) int
	minInt = func (a, b int) int {
		if a > b {
			return b
		} else {
			return a
		}
	}

	dp := make([]int, len(data))
	dp[0] = data[0][0]

	for i := 1; i < len(data); i++ {
		n := len(data[i])
		last := dp[0]
		dp[0] = dp[0] + data[i][0]

		for k := 1; k < n - 1; k++ {
			tmp := dp[k]

			min := minInt(last, dp[k])
			dp[k] = data[i][k] + min

			last = tmp
		}

		dp[n - 1] = last + data[i][n - 1]
	}

	min := dp[0]

	for _, v := range dp {
		if min > v {
			min = v
		}
	}

	return min
}

func MinimumTotal() {
	data := [][]int {
		[]int {2},
		[]int {3, 4},
		[]int {6, 5, 7},
		[]int {4, 1, 8, 3},
	}

	fmt.Printf("<120> ")
	fmt.Println(minimumTotal(data))
}

