package problems

import "fmt"

func convertToTitle(n int) string {
	if n == 0 {
		return ""
	}

	out := ""

	for n != 0 {
		n--
		d := n % 26

		out = string('A' + d) + out
		n = n / 26
	}

	return out
}

func ConvertToTitle() {
	fmt.Printf("<168> ")
	fmt.Println(convertToTitle(26))
}

