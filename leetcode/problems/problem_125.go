package problems

import "fmt"

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	d := []byte(s)

	var isLetter func (a byte) bool
	isLetter = func (a byte) bool {
		if a >= 'a' && a <= 'z' {
			return true
		} else if a >= 'A' && a <= 'Z' {
			return true
		} else if a >= '0' && a <= '9' {
			return true
		} else {
			return false
		}
	}

	var isSame func (a, b byte) bool
	isSame = func(a, b byte) bool {
		if a == b {
			return true
		} else if a >= '0' && a <= '9' || b >= '0' && b <= '9' {
			return false
		} else if a - b == 'a' - 'A' || b - a == 'a' - 'A' {
			return true
		} else {
			return false
		}
	}

	i, j := 0, len(d) - 1

	for i <= j {
		if !isLetter(d[i]) {
			i++
			continue
		}

		if !isLetter(d[j]) {
			j--
			continue
		}

		if isSame(d[i], d[j]) {
			i++
			j--
		} else {
			return false
		}
	}

	return true
}

func IsPalindrome() {
	fmt.Printf("<125> ")
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

