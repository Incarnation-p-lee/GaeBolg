package problems

import "fmt"
import "math"

func maximumGap(nums []int) int {
	n := nums

	if len(nums) < 2 {
		return 0
	}

	max, min := math.MinInt32, math.MaxInt32

	for _, v := range n {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	step := (max - min + 1) / len(n) + 1
	bmax := make([]int, len(n) + 1)
	bmin := make([]int, len(n) + 1)

	for _, v := range n {
		i := (v - min) / step
		if bmax[i] == 0 {
			bmax[i], bmin[i] = v, v
			continue
		}

		if v > bmax[i] {
			bmax[i] = v
		}

		if v < bmin[i] {
			bmin[i] = v
		}
	}

	out, pmax := 0, bmax[0]

	for i, v := range bmax {
		if bmin[i] != 0 {
			d := bmin[i] - pmax
			if d > out {
				out = d
			}

			pmax = v
		}
	}

	return out
}

func MaximumGap() {
	fmt.Printf("<164> ")
	fmt.Println(maximumGap([]int {1, 9, 3, 5, 8}))
}

