package problems

import "fmt"

func plusOne(digits []int) []int {
	d := digits

	if len(d) == 0 {
		return []int{1}
	}

	isCarry := false

	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits) - 1 || isCarry {
			n := d[i] + 1
			d[i] = n % 10

			if n > 9 {
				isCarry = true
			} else {
				isCarry = false
			}
		} else {
			break
		}
	}

	if isCarry {
		d = append([]int{1}, d...)
	}

	return d
}

func PlusOne() {
	fmt.Printf("<066> ")
	fmt.Println(plusOne([]int{1, 2, 9}))
}

