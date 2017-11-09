package problems

import "fmt"

func multiply(num1 string, num2 string) string {
	var out []byte
	var addstring func (n1 []byte, n2 []byte) []byte
	var singleMultiply func (n []byte, b byte, sft int) []byte

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	singleMultiply = func (n []byte, b byte, sft int) []byte {
		var result []byte

		if b == '0' {
			return []byte{'0'}
		} else if b == '1' {
			result = make([]byte, len(n))
			copy(result, n)
		} else {
			carry := 0
	
			for i := len(n) - 1; i >= 0; i-- {
				v := n[i]
				d := int(v - '0') * int(b - '0') + carry
				carry = d / 10
				d = d % 10

				result = append([]byte{byte(d) + '0'}, result...)
			}
	
			if carry > 0 {
				result = append([]byte{byte(carry) + '0'}, result...)
			}
		}

		for i := 0; i < sft; i++ {
			result = append(result, '0')
		}

		return result
	}

	addstring = func (n1 []byte, n2 []byte) []byte {
		var result []byte

		carry, i, j := 0, len(n1) - 1, len(n2) - 1

		for i >= 0 && j >= 0 {
			d := int(n1[i] - '0') + int(n2[j] - '0') + carry
			carry = d / 10
			d = d % 10

			result = append([]byte{byte(d) + '0'}, result...)
			i--
			j--
		}

		var addAppend func (n []byte, k int)
		addAppend = func (n []byte, k int) {
			for i := k; i >= 0; i-- {
				d := int(n[i] - '0') + carry
				carry = d / 10
				d = d % 10

				result = append([]byte{byte(d) + '0'}, result...)
			}
		}

		if i >= 0 {
			addAppend(n1, i)
		}

		if j >= 0 {
			addAppend(n2, j)
		}

		if carry == 1 {
			result = append([]byte{'1'}, result...)
		}

		return result
	}

	n1, n2 := []byte(num1), []byte(num2)

	for i := len(n1) - 1; i >= 0; i-- {
		tmp := singleMultiply(n2, n1[i], len(n1) - 1 - i)
		out = addstring(out, tmp)
	}

	return string(out)
}

func Multiply() {
	fmt.Printf("<043> ")
	fmt.Println(multiply("214", "234"))
}

