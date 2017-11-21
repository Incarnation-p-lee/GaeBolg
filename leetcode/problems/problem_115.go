package problems

import "fmt"

func numDistinct(s string, t string) int {
	count := 0
	bs, bt := []byte(s), []byte(t)

	var numDistinctRecursive func (s, t []byte, out *int)
	numDistinctRecursive = func (s, t []byte, out *int) {
		if len(t) == 0 {
			*out += 1
		} else if len(t) <= len(s) {
			for i, v := range s {
				if v == t[0] {
					numDistinctRecursive(s[i + 1:], t[1:], out)
				}
			}
		}
	}

	numDistinctRecursive(bs, bt, &count)

	return count
}

func numDistinctDp(s string, t string) int {
	if len(t) > len(s) {
		return 0
	}

	bs, bt := []byte(s), []byte(t)

	m, n := len(bt), len(bs)
	dp := make([][]int, m)

	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	if bt[0] == bs[0] {
		dp[0][0] = 1
	}

	for j := 1; j < n; j++ {
		if bs[j] == bt[0] {
			dp[0][j] = dp[0][j - 1] + 1
		} else {
			dp[0][j] = dp[0][j - 1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if i > j {
				dp[i][j] = 0
			} else if bt[i] == bs[j] {
				dp[i][j] = dp[i][j - 1] + dp[i - 1][j - 1]
			} else {
				dp[i][j] = dp[i][j - 1]
			}
		}
	}

	return dp[m - 1][n - 1]
}

func NumDistinct() {
	fmt.Printf("<115> ")
	fmt.Println(numDistinct("rabbbit", "rabbit"))

	fmt.Printf("<115> ")
	fmt.Println(numDistinctDp("rabbbit", "rabbit"))
}

