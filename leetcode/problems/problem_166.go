package problems

import "fmt"
import "math"
import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
	n := int64(numerator)
	d := int64(denominator)

	if d == 0 {
		return ""
	} else if n == 0 {
		return "0"
	} else if n == math.MinInt32 && d == 1 {
		return "-2147483648"
	} else if n == math.MinInt32 && d == -1 {
		return "2147483648"
	}

	rmap := make(map[int]int)
	out := ""
	reminder := 0
	mantiass := make([]byte, 0)
	isMinus := false

	if n < 0 {
		n = -n
		isMinus = true
	}

	if d < 0 {
		d = -d
		isMinus = !isMinus
	}

	if n < d {
		out += "0."
		reminder = int(n)
	} else {
		out += strconv.Itoa(int(n / d))
		reminder = int(n % d)

		if reminder == 0 {
			return out
		}

		out += "."
	}

	for {
		if reminder == 0 {
			break
		}

		if v, ok := rmap[reminder]; ok {
			idx := v
			left := make([]byte, len(mantiass[:idx]))
			copy(left, mantiass[:idx])
			right := make([]byte, len(mantiass[idx:]))
			copy(right, mantiass[idx:])

			mantiass = append(left, '(')
			mantiass = append(mantiass, right...)
			mantiass = append(mantiass, ')')
			break
		}

		rmap[reminder] = len(mantiass)
		num := reminder * 10

		if int64(num) < d {
			mantiass = append(mantiass, '0')
			reminder = num
		} else {
			v := int64(num) / d
			mantiass = append(mantiass, '0' + byte(v))
			reminder = int(int64(num) % d)
		}
	}

	out += string(mantiass)

	if isMinus {
		out = "-" + out
	}

	return out
}

func FractionToDecimal() {
	fmt.Printf("<166> ")
	fmt.Println(fractionToDecimal(1, 6))
}

