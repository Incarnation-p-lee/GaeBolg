package problems

import "fmt"

func isValid(s string) bool {
	var stack []byte

	data := []byte(s)

	for _, v := range(data) {
		if len(stack) == 0 {
			stack = append(stack, v)
			continue
		}

		d := stack[len(stack) - 1]

		switch (d) {
			case '(':
			if v == ')' {
				stack = stack[:len(stack) - 1]
			} else if v == ']' || v == '}' {
				return false
			} else {
				stack = append(stack, v)
			}
			case '[':
			if v == ']' {
				stack = stack[:len(stack) - 1]
			} else if v == ')' || v == '}' {
				return false
			} else {
				stack = append(stack, v)
			}
			case '{':
			if v == '}' {
				stack = stack[:len(stack) - 1]
			} else if v == ')' || v == ']' {
				return false
			} else {
				stack = append(stack, v)
			}
			default:
			return false
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

func IsValid() {
	fmt.Printf("<020> ")
	fmt.Println(isValid("()[]{}"))
}

