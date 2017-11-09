package problems

import "fmt"

func myPow(x float64, n int) float64 {
	v := n

	if v < 0 {
		v = -v
	}

	k, out := uint32(v), float64(1.0)

	if v == 0 {
		return float64(1.0)
	} else if v == 1 {
		out = x
	} else {
		dp := make([]float64, 1)
		dp = append(dp, x)

		for i := 2; k != 0; i++ {
			d := dp[i - 1]
			dp = append(dp, d * d)

			if k & 0x1 == 0x1 {
				out = out * dp[i - 1]
			}

			k = k >> 1
		}
	}

	if n < 0 {
		out = 1.0 / out
	}

	return out
}

func MyPow() {
	fmt.Printf("<050> ")
	fmt.Println(myPow(2.3, 10))
}

