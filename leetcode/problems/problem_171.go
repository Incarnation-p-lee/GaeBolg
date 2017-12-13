package problems

import "fmt"

func titleToNumber(s string) int {
	if len(s) == 0 {
		return 0
	}

	out, d := 0, []byte(s)

	for _, v := range d {
		out = out * 26 + int(v - 'A') + 1
	}

	return out
}

func TitleToNumber() {
	fmt.Printf("<171> ")
	fmt.Println(titleToNumber("AZZ"))
}

