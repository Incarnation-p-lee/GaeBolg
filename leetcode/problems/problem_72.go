package problems

import "fmt"

func minDistance(word1 string, word2 string) int {
	d1 := []byte(word1)
	d2 := []byte(word2)

	if len(d1) == 0 {
		return len(d2)
	} else if len(d2) == 0 {
		return len(d1)
	}

	var minInt func (a, b int) int
	minInt = func (a, b int) int {
		if a > b {
			return b
		} else {
			return a
		}
	}

	m, n := len(d1) + 1, len(d2) + 1
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for k := 0; k < n; k++ {
		dp[0][k] = k
	}

	for k := 0; k < m; k++ {
		dp[k][0] = k
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			left, up, last := dp[i][j - 1], dp[i - 1][j], dp[i - 1][j - 1]
			min := minInt(left + 1, up + 1)

			if d1[i - 1] != d2[j - 1] {
				last = last + 1
			}

			dp[i][j] = minInt(last, min)
		}
	}

	return dp[m - 1][n - 1]
}

func MinDistance() {
	fmt.Printf("<072> ")
	fmt.Println(minDistance("abed", "adbdabd"))
}

