package problems

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	b := []byte(s)
	w := wordDict

	if len(b) == 0 || len(w) == 0 {
		return false
	}

	smap := make(map[string]int)

	for _, v := range w {
		smap[v] = 1
	}

	dp := make([]bool, len(b) + 1)
	dp[0] = true

	for i := 0; i < len(b); i++ {
		canBreak := false
		for k := 0; k <= i; k++ {
			_, ok := smap[string(b[k:i + 1])]
			canBreak = canBreak || (dp[k] && ok)

			if canBreak {
				break
			}
		}

		dp[i + 1] = canBreak
	}

	return dp[len(b)]
}

func WordBreak() {
	s := "a"
	w := []string {"a", "code"}

	fmt.Printf("<139> ")
	fmt.Println(wordBreak(s, w))
}

