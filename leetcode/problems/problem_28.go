package problems

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return -1
	} else if len(haystack) == 0 {
		return 0
	}

	loc := -1

	h, n := []byte(haystack), []byte(needle)

	for i, v := range h {
		if v == n[0] {
			j := i
			isMatched := true

			for _, d := range n {
				if j >= len(h) || h[j] != d {
					isMatched = false
					break
				}

				j++
			}

			if isMatched {
				loc = i
				break
			}
		}
	}

	return loc
}

func StrStr() {
	fmt.Printf("<028> ")
	fmt.Println(strStr("hello", "ll"))
}

