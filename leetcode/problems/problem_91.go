package problems

import "fmt"
import "strconv"

func numDecodings(s string) int {
	n := len(s)
	data := []byte(s)

	if len(data) == 0 || data[0] == '0' {
		return 0
	}

	dp0, dp1, dp2 := 1, 1, 1

	for i := 1; i < n; i++ {
		dp := 0
		if data[i] == '0' {
			if data[i - 1] == '2' || data[i - 1] == '1' {
				dp = dp1
			} else {
				return 0
			}
		} else {
			d, _ := strconv.Atoi(string(data[i - 1:i + 1]))

			if d > 26 {
				dp = dp2
			} else if d < 10 {
				if data[i - 2] == '2' || data[i - 2] == '1' {
					dp = dp0
				} else {
					return 0
				}
			} else {
				dp = dp2 + dp1
			}
		}

		dp0, dp1, dp2 = dp1, dp2, dp
	}

	return dp2
}

func NumDecodings() {
	fmt.Printf("<091> ")
	fmt.Println(numDecodings("10223101"))
}

