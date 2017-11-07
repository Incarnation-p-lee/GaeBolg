package problems

import "fmt"

func longestValidParentheses(s string) int {
	var stack []int

	if len(s) == 0 {
		return 0
	}

	l := 0
	cnt := 0
	data := []byte(s)

	for i, v := range data {
		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			k := stack[len(stack) - 1]
			top := data[k]

			switch top {
			case '(':
				if v == '(' {
					stack = append(stack, i)
				} else {
					p := -1
					stack = stack[:len(stack) - 1]

					if len(stack) > 0 {
						p = stack[len(stack) - 1]
					}

					cnt = i - p
					if cnt > l {
						l = cnt
					}
				}
			case ')':
				stack = append(stack, i)
			}
		}
	}

	return l
}

func LongestValidParentheses() {
	fmt.Printf("<032> ")
	fmt.Println(longestValidParentheses("()"))
}

