package problems

import "fmt"

func largestRectangleArea(heights []int) int {
	h := heights

	if len(h) == 0 {
		return 0
	}

	var expandArea func (k int) int
	expandArea = func (k int) int {
		d := h[k]
		l, r := k, k

		for ; l >= 0; l-- {
			if h[l] < d {
				break
			}
		}

		l++

		for ; r < len(h); r++ {
			if h[r] < d {
				break
			}
		}

		r--

		return (r - l + 1) * d
	}

	max := -1

	for i, _ := range h {
		d := expandArea(i)
		if d > max {
			max = d
		}
	}

	return max
}

func largestRectangleArea2(heights []int) int {
	h := heights

	if len(h) == 0 {
		return 0
	}

	i, max := 0, -1
	h = append(h, 0)
	stack := make([]int, 0)

	for i < len(h) {
		if len(stack) == 0 || h[stack[len(stack) - 1]] < h[i]  {
			stack = append(stack, i)
			i++
		} else {
			d := 0
			top := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]

			if len(stack) == 0 {
				d = i
			} else {
				d = i - 1 - stack[len(stack) - 1]
			}

			area := d * h[top]

			if area > max {
				max = area
			}
		}
	}

	return max
}

func LargestRectangleArea() {
	data := []int {2, 1, 5, 6, 2, 3}

	fmt.Printf("<084> ")
	fmt.Println(largestRectangleArea(data));

	fmt.Printf("<084> ")
	fmt.Println(largestRectangleArea2(data));
}

