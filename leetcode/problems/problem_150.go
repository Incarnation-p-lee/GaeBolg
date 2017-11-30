package problems

import "fmt"
import "strconv"

func evalRPN(tokens []string) int {
	t := tokens

	if len(t) == 0 {
		return 0
	}

	stack := make([]string, 0)

	for _, v := range t {
		if v != "+" && v != "-" && v != "*" && v != "/" {
			stack = append(stack, v)
		} else {
			sa, sb := stack[len(stack) - 1], stack[len(stack) - 2]
			stack = stack[:len(stack) - 2]
			a, _ := strconv.Atoi(sa)
			b, _ := strconv.Atoi(sb)

			switch v {
			case "+":
				stack = append(stack, strconv.Itoa(b + a))
			case "-":
				stack = append(stack, strconv.Itoa(b - a))
			case "*":
				stack = append(stack, strconv.Itoa(b * a))
			case "/":
				stack = append(stack, strconv.Itoa(b / a))
			}
		}
	}

	val, _ := strconv.Atoi(stack[0])

	return val
}

func EvalRPN() {
	fmt.Printf("<150> ")
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
}

