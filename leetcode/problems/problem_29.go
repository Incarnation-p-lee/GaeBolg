package problems

import "fmt"
import "math"

func divide(dividend int, divisor int) int {

	if divisor == 0 {
		return math.MaxInt32
	} else if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	} else if dividend == math.MinInt32 && divisor == -2 {
		return math.MaxInt32 / 2 + 1
	}

	out := 0
	isMinus := false

	if dividend < 0 {
		isMinus = !isMinus
		dividend = -dividend
	}

	if divisor < 0 {
		isMinus = !isMinus
		divisor = -divisor
	}

	for d := divisor; dividend >= d; d = divisor {
		s := 1

		for {
			if dividend < d + d {
				break
			}
			d += d
			s += s
		}


		dividend = dividend - d
		out += s
	}

	if isMinus {
		out = -out
	}

	return out
}

func Divide() {
	fmt.Printf("<029> ")
	fmt.Println(divide(-2147483648, -1))
}

