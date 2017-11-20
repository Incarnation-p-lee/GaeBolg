package problems

import "fmt"

func isInterleaveRecursive(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1) + len(s2) {
		return false
	}

	b1, b2, b3 := []byte(s1), []byte(s2), []byte(s3)

	var isInter func (d1, d2, d3 []byte) bool
	isInter = func (d1, d2, d3 []byte) bool {
		if len(d1) == 0 && len(d2) == 0 && len(d3) == 0 {
			return true
		}

		isMatch := false

		if len(d1) > 0 && d1[0] == d3[0] {
			isMatch = isInter(d1[1:], d2, d3[1:])
		}

		if !isMatch && len(d2) > 0 && d2[0] == d3[0] {
			isMatch = isInter(d1, d2[1:], d3[1:])
		}

		return isMatch
	}

	return isInter(b1, b2, b3)
}

/*
 * s1 = "ab", s2 = "aa", s3 = "abaa"
 *       | a | a |
 *   +---+---+---+
 *   | 0 | 1 | 0 |
 * --+---+---+---+
 * a | 1 | 0 | 0 |
 * --+---+---+---+
 * b | 1 | 1 | 1 |
 * --+---+---+---+
 */
func isInterleaveDp(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1) + len(s2) {
		return false
	}

	b1, b2, b3 := []byte(s1), []byte(s2), []byte(s3)
	m, n := len(b1) + 1, len(b2) + 1

	dp := make([][]bool, m)
	for i, _ := range dp {
		dp[i] = make([]bool, n)
	}

	dp[0][0] = true

	for i, v := range b1 {
		if v == b3[i] {
			dp[i + 1][0] = true
		} else {
			break
		}
	}

	for j, v := range b2 {
		if v == b3[j] {
			dp[0][j + 1] = true
		} else {
			break
		}
	}

	for i := 1; i <= len(b1); i++ {
		vb1 := b1[i - 1]
		for j := 1; j <= len(b2); j++ {
			vb2 := b2[j - 1]
			k := i + j
			vb3 := b3[k - 1]

			if vb3 == vb1 {
				dp[i][j] = dp[i - 1][j] || dp[i][j]
			}

			if vb3 == vb2 {
				dp[i][j] = dp[i][j - 1] || dp[i][j] 
			}
		}
	}

	return dp[m - 1][n - 1]
}

func IsInterleave() {
	fmt.Printf("<097> ")
	fmt.Println(isInterleaveRecursive("aa", "ab", "abaa"))

	fmt.Printf("<097> ")
	fmt.Println(isInterleaveDp("aa", "ab", "aaab"))
}

