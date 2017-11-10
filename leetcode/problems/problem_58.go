package problems

import "fmt"

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	data, i, l := []byte(s), 0, 0

	for i < len(data) {
		if data[i] == ' ' {
			i++
		} else {
			l = 0

			for i < len(data) && data[i] != ' ' {
				l++
				i++
			}
		}
	}

	return l
}

func LengthOfLastWord() {
	fmt.Printf("<058> ")
	fmt.Println(lengthOfLastWord("Hello World  "))
}

