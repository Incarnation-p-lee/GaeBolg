package problems

import "fmt"
import "strconv"

func numDecodings(s string) int {
	n := len(s)
	data := []byte(s)
	dp := make([]int, n + 1)

	if len(s) == 0 {
		return 0
	} else if len(s) == 1 && s == "0" {
		return 0
	} else if len(s) == 1 {
		return 1
	}

	dp[0] = 1
	d, _ := strconv.Atoi(string(data[:2]))

	if data[0] == '0' {
		return 0
	} else {
		dp[1] = 1
	}

	if d < 10 {
		return 0
	} else if data[1] == '0' {
		if d == 10 || d == 20 {
			dp[2] = 1
		} else {
			return 0
		}
	} else if d > 26 {
		dp[2] = 1
	} else {
		dp[2] = 2
	}

	for i := 2; i < n; i++ {
		if data[i] == '0' {
			if data[i - 1] <= '2' && data[i - 1] > '0' {
				dp[i + 1] = dp[i - 1]
			} else {
				return 0
			}
		} else {
			d, _ := strconv.Atoi(string(data[i - 1:i + 1]))

			if d > 26 {
				dp[i + 1] = dp[i]
			} else if d < 10 {
				if data[i - 2] <= '2' && data[i - 2] > '0' {
					dp[i + 1] = dp[i - 2]
				} else {
					return 0
				}
			} else {
				dp[i + 1] = dp[i] + dp[i - 1]
			}
		}
	}

	return dp[n]
}

func NumDecodings() {
	fmt.Printf("<091> ")
	fmt.Println(numDecodings("10223101"))
}

