package problems

import "fmt"

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var maxInt func (a, b int) int
	maxInt = func (a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	var minInt func (a, b int) int
	minInt = func (a, b int) int {
		if a > b {
			return b
		} else {
			return a
		}
	}

	n := nums
	out, min, max := n[0], n[0], n[0]

	for i := 1; i < len(n); i++ {
		a := min * n[i]
		b := max * n[i]
		c := n[i]

		min = minInt(minInt(a, b), c)
		max = maxInt(maxInt(a, b), c)

		if max > out {
			out = max
		}
	}

	return out
}

func MaxProduct() {
	data := []int {-4, -3, -2}

	fmt.Printf("<152> ")
	fmt.Println(maxProduct(data))
}

