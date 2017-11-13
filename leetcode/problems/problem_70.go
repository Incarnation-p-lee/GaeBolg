package problems

import "fmt"

func climbStairs(n int) int {
	dp := make([]int, n + 3)

	dp[0] = 0
	dp[1] = 1
	dp[2] = 2

	i := 3

	for i <= n {
		dp[i] = dp[i - 1] + dp[i - 2]
		i++
	}

	return dp[n]
}

func ClimbStairs() {
	fmt.Printf("<070> ")
	fmt.Println(climbStairs(8))
}

