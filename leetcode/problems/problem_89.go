package problems

import "fmt"

func grayCode(n int) []int {
	var out []int    

	if n == 0 {
		return []int{0}
	} else if n == 1 {
		return []int{0, 1}
	}

	out = append(out, 0, 1)

	for i := 1; i < n; i++ {
		sz := len(out)

		for j := sz - 1; j >= 0; j-- {
			d := uint32(out[j]) + (uint32(1) << uint32(i))
			out = append(out, int(d))
		}
	}

	return out
}

func GrayCode() {
	n := 2

	fmt.Printf("<089> ")
	fmt.Println(grayCode(n))
}

