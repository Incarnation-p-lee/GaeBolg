package problems

import "fmt"

func calculateMinimumHP(dungeon [][]int) int {
	d := dungeon

	if len(d) == 0 || len(d[0]) == 0 {
		return 0
	}

	n, m := len(d), len(d[0])
	dp := make([][]int, n)

	var maxInt func (a, b int) int
	maxInt = func (a, b int) int {
		if a < b {
			return b
		} else {
			return a
		}
	}

	var minInt func (a, b int) int
	minInt = func (a, b int) int {
		if a > b {
			return b
		} else {
			return a
		}
	}

	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	dp[n - 1][m - 1] = maxInt(1, 1 - d[n - 1][m - 1])

	for i := n - 2; i >= 0; i-- {
		dp[i][m - 1] = maxInt(1, dp[i + 1][m - 1] - d[i][m - 1])
	}

	for j := m - 2; j >= 0; j-- {
		dp[n - 1][j] = maxInt(1, dp[n - 1][j + 1] - d[n - 1][j])
	}

	for i := n - 2; i >= 0; i-- {
		for j := m - 2; j >= 0; j-- {
			v := minInt(dp[i + 1][j], dp[i][j + 1]) - d[i][j]
			dp[i][j] = maxInt(1, v)
		}
	}

	return dp[0][0]
}

func CalculateMinimumHP() {
	data := [][]int {
		[]int {-2,  -3,  3},
		[]int {-5, -10,  1},
		[]int {10,  30, -5},
	}

	fmt.Printf("<174> ")
	fmt.Println(calculateMinimumHP(data))
}

