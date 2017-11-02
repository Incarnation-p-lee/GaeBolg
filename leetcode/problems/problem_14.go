package problems

import "fmt"

func longestCommonPrefix(strs []string) string {
	var i int
	var is_done bool
	var data [][]byte

	if len(strs) == 0 {
		return ""
	}

	for _, v := range strs {
		data = append(data, []byte(v))
	}

	for i = 0; i < len(data[0]); i++ {
		d := data[0][i]
		for _, v := range data {
			if i >= len(v) {
				is_done = true
				break
			}

			if d != v[i] {
				is_done = true
				break
			}
		}

		if is_done {
			break
		}
	}

	return string(data[0][:i])
}

func LongestCommonPrefix() {
	var s []string = []string {"abcd", "abcd", "abcd", "abcd"}

	fmt.Printf("<014> ")
	fmt.Println(longestCommonPrefix(s))
}

