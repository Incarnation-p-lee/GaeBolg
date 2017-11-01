package problems

import "fmt"

func myAtoi(str string) int {
	var d int64
	var is_minus, is_set, is_num bool
	data := []byte(str)

	is_set, is_num = false, false

	for _, v := range data {
		if v == ' ' && !is_num {
			continue
		} else if (v == '-' || v == '+') && is_set {
			return 0
		} else if v == '-' {
			is_minus = true
			is_set = true
			is_num = true
		} else if v == '+' {
			is_minus = false
			is_set = true
			is_num = true
		} else if v >= '0' && v <= '9' {
			is_num = true
			d = d * 10 + int64(v - '0')

			if is_minus && d > 2147483648 {
				return -2147483648
			} else if !is_minus && d > 2147483647 {
				return 2147483647
			}
		} else {
			break
		}
	}

	if is_minus {
		d = -d
	}

	return int(d)
}

func MyAtoi() {
	fmt.Printf("<008> ") 
	fmt.Println(myAtoi("   + 0 123"))
}

