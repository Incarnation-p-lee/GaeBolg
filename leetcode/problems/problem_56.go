package problems

import "fmt"

func merge(itl []Interval) []Interval {
	if len(itl) == 0 {
		return []Interval{}
	}

	out := make([]Interval, 0)
	max := 0

	for _, v := range itl {
		if v.End > max {
			max = v.End
		}
	}

	array := make([]int, max + 2)

	for _, v := range itl {
		s, e := v.Start, v.End

		if s == e && array[s] == 0 {
			array[s] = 2
		} else {
			for s < e {
				array[s] = 1
				s++
			}
		}
	}

	for i := 0; i < len(array); i++ {
		s, e := 0, 0

		if array[i] == 2 {
			out = append(out, Interval{i, i})
		} else if array[i] == 1 {
			s = i

			for array[i] == 1 {
				i++
			}

			e = i

			out = append(out, Interval{s, e})
		}
	}

	return out
}

func Merge() {
	data := []Interval{{2, 3}, {4, 5}, {6, 7}, {1, 10}, {11, 12}, {12, 12}}

	fmt.Printf("<056> ")
	fmt.Println(merge(data))
}

