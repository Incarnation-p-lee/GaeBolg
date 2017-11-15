package problems

import "fmt"
import "math"

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	hash := make([]int, 128)

	for _, k := range []byte(t) {
		hash[k]++
	}

	data := []byte(s)
	l, r, count := 0, 0, len(t)
	min, minl, minr := math.MaxInt32, 0, 0

	for ; r < len(data); r++ {
		v := data[r]

		hash[v]--

		if hash[v] > -1 {
			count--
		}

		for ; count == 0; l++ { /* Get candidate window, then try to narrow down */
			if r - l + 1 < min {
				min = r - l + 1
				minl, minr = l, r + 1
			}

			v := data[l]
			hash[v]++

			if hash[v] > 0 {
				count++
			}
		}
	}

	if min == math.MaxInt32 {
		return ""
	}

	return string(data[minl:minr])
}

func MinWindow() {
	fmt.Printf("<076> ")
	fmt.Println(minWindow("bba", "ab"))
}

