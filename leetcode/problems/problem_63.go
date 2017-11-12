package problems

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	d := obstacleGrid

	if len(d) == 0 {
		return 0
	}

	isBlocked := false
	dp := make([]int, len(d[0]))

	for i, v := range d[0] {
		if v == 1 {
			isBlocked = true
		}

		if !isBlocked {
			dp[i] = 1
		} else {
			dp[i] = 0
		}
	}

	for i := 1; i < len(d); i++ {
		for j, v := range d[i] {
			if v == 1 {
				dp[j] = 0
			} else if j != 0 {
				dp[j] = dp[j - 1] + dp[j]
			}
		}
	}

	return dp[len(d[0]) - 1]
}

func UniquePathsWithObstacles() {
	data := [][]int {
		[]int {0, 0, 0},
		[]int {0, 1, 0},
		[]int {0, 0, 0},
	}

	fmt.Printf("<063> ")
	fmt.Println(uniquePathsWithObstacles(data))
}

