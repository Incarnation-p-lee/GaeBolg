package problems

import "fmt"

func insert(intervals []Interval, newInterval Interval) []Interval {
	if len(intervals) == 0 {
		return []Interval{newInterval}
	}

	new := newInterval
	tmp := make([]Interval, 0)
	left, right := -1, -1

	for _, v := range intervals {
		if new.Start >= v.Start && new.Start <= v.End {
			left = v.Start
		}

		if new.End >= v.Start && new.End <= v.End {
			right = v.End
		}

		if v.End < new.Start {
			tmp = append(tmp, v)
		}

		if v.Start > new.End {
			tmp = append(tmp, v)
		}
	}

	if left == -1 {
		left = new.Start
	}

	if right == -1 {
		right = new.End
	}

	out := make([]Interval, 0)

	if len(tmp) == 0 {
		out = append(out, Interval{left, right})
	} else if right < tmp[0].Start {
		out = append([]Interval{{left, right}}, tmp...)
	} else if left > tmp[len(tmp) - 1].End {
		out = append(tmp, Interval{left, right})
	} else {
		inserted := false

		for _, v := range tmp {
			if v.Start > left && !inserted {
				inserted = true
				out = append(out, Interval{left, right})
			}

			out = append(out, v)
		}
	}

	return out
}

func Insert() {
	data := []Interval{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}

	fmt.Printf("<057> ")
	fmt.Println(insert(data, Interval{1, 10}))
}

