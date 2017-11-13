package problems

import "fmt"
import "strings"

func simplifyPath(path string) string {
	data := []byte(path)

	if len(data) == 0 {
		return path
	}

	pcount := 0
	stack := make([]byte, 0)
	stack = append(stack, data[0])

	var popData func ()
	popData = func () {
		if string(stack) == "/.." || string(stack) == "/../" {
			stack = stack[:1]
			return
		}

		if len(stack) > 3 && string(stack[len(stack) - 3:]) == "/.." {
			stack = stack[:len(stack) - 3]
		} else if len(stack) > 4 && string(stack[len(stack) - 4:]) == "/../" {
			stack = stack[:len(stack) - 4]
		} else {
			return
		}

		for stack[len(stack) - 1] != '/' {
			stack = stack[:len(stack) - 1]
		}
	}

	for i := 1; i < len(data); i++ {
		v := data[i]
		top := stack[len(stack) - 1]

		if v == '.' {
			pcount++
		}

		switch top {
		case '.':
			if v == '/' && pcount == 1 {
				stack = stack[:len(stack) - 1]
			} else if v == '/' && pcount == 2 {
				popData()
			} else {
				stack = append(stack, v)
			}
		case '/':
			if v != '/' {
				stack = append(stack, v)
			}
		default:
			stack = append(stack, v)
		}

		if v != '.' {
			pcount = 0
		}
	}

	popData()

	if len(stack) > 2 && string(stack[len(stack) - 2:]) == "/." {
		stack = stack[:len(stack) - 2]
	} else if len(stack) == 2 && string(stack[len(stack) - 2:]) == "/." {
		stack = stack[:len(stack) - 1]
	}

	if len(stack) > 1 && stack[len(stack) - 1] == '/' {
		stack = stack[:len(stack) - 1]
	}

	return string(stack)
}

func simplifyPath2(path string) string {
	if len(path) == 0{
		return path
	}

	data := make([]string, 0)

	for _, v := range strings.Split(path, "/") {
		if v == ".." && len(data) > 0 {
			data = data[:len(data) - 1]
		} else if v != "." && len(v) > 0 && v != ".." {
			data = append(data, v)
		}
	}

	return "/" + strings.Join(data, "/")
}

func SimplifyPath() {
	fmt.Printf("<071> ")
	fmt.Println(simplifyPath("/../abdc/.../ad/.."))

	fmt.Printf("<071> ")
	fmt.Println(simplifyPath2("/../abdc/.../ad/.."))
}

