package problems

import "fmt"

func numTrees(n int) int {
	dp := make([]int, n + 1)

	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i - 1 - j]
		}
	}

	return dp[n]
}

func NumTrees() {
	fmt.Printf("<096> ")
	fmt.Println(numTrees(3))
}

