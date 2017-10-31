package problems

import "fmt"

func longestPalindrome(s string) string {
	data := []byte(s)
	var out string
	var m, n, k, length int

	for i := 0; i < len(data); {
		d := data[i]
		j := i

		for ; j < len(data); j++ {
			if d != data[j] {
				break
			}
		}

		k, n  = j, j

		if i > 0 {
			m = i - 1
			for m >= 0 && n < len(data) {
				if data[m] != data[n] {
					break
				}
				m--
				n++
			}
		} else {
			m = -1
		}

		if n - m - 1 > length {
			length = n - m - 1
			out = string(data[m + 1: n])
		}

		i = k
	}

	return out
}

func LongestPalindrome() {
	s := "babad"

	fmt.Printf("<005> ")
	fmt.Println(longestPalindrome(s))
}

