package problems

import "fmt"

func trapInternal(h []int, i, j int, out *int) {
	if j - i < 2 {
		return
	}

	var dataTrap func (l, r int) int
	dataTrap = func (l, r int) int {
		if r - l < 2 {
			return 0
		}

		data := 0
		limit := h[l]

		if h[l] > h[r] {
			limit = h[r]
		}

		for k := l + 1; k < r; k++ {
			data += limit - h[k]
		}

		return data
	}

	max, second := -1, -1
	mi, si := -1, -1
	l, r := i, j

	for ; l <= r; l++ {
		if h[l] > max {
			if max != -1 {
				second = max
				si = mi
			}

			max = h[l]
			mi = l
		}

		if h[l] > second && h[l] <= max && l != mi {
			second = h[l]
			si = l
		}
	}

	if mi > si {
		l, r = si, mi
	} else {
		l, r = mi, si
	}

	*out += dataTrap(l, r)

	trapInternal(h, i, l, out)
	trapInternal(h, r, j, out)
}

func trap(height []int) int {
	var out = 0

	trapInternal(height, 0, len(height) - 1, &out)

	return out
}

func Trap() {
	height := []int{0, 1, 0, 2, 1, 0, 7, 3, 2, 1, 2, 1}

	fmt.Printf("<042> ")
	fmt.Println(trap(height))
}

