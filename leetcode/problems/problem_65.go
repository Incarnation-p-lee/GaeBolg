package problems

import "fmt"

func isNumber(s string) bool {
	data := []byte(s)
	l, r := 0, 0

	for i, v := range data {
		if v != ' ' {
			l = i
			break
		}
	}

	for i := len(data) - 1; i >= 0; i-- {
		if data[i] != ' ' {
			r = i
			break
		}
	}

	data = data[l: r + 1]

	if len(data) == 0 {
		return false
	}

	last := byte(0)
	isPoint, isE, hasNumber := false, false, false

	for i, v := range data {
		if (v == '-' || v == '+') && (i == 0 || last == 'e') {
			;
		} else if v == ' ' {
			return false
		} else if v == '.' {
			if isPoint || last == 'e' || isE {
				return false
			} else {
				isPoint = true
			}
		} else if v == 'e' {
			if isE || i == 0 || !hasNumber {
				return false
			} else {
				isE = true
			}
		} else if v < '0' || v > '9' {
			return false
		} else {
			hasNumber = true
		}

		last = v
	}

	if last == 'e' || last == '-' || last == '+' || !hasNumber {
		return false
	}

	return true
}

func IsNumber() {
	fmt.Printf("<065> ")
	fmt.Println(isNumber(".1"))
}

