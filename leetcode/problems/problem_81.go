package problems

import "fmt"

func search2(nums []int, target int) bool {
	d, t := nums, target

	if len(d) == 0 {
		return false
	}

	var bFind func (i, j, t int) bool
	bFind = func (i, j, t int) bool {
		for i <= j {
			m := (i + j) / 2

			if d[m] == t {
				return true
			}

			if t < d[m] {
				j = m - 1
			} else if t > d[m] {
				i = m + 1
			}
		}

		return false
	}

	l, r := 0, len(d) - 1

	for l <= r {
		m := (l + r) / 2
		med := d[m]

		if med == t {
			return true
		}

		if d[l] == med && d[r] == med {
			l++
			r--
			continue
		}

		if d[l] <= med { /* sorted */
			if t < med && t >= d[l] {
				return bFind(l, m, t)
			} else {
				l = m + 1
				continue
			}
		}

		if med <= d[r] { /* sorted */
			if t > med && t <= d[r] {
				return bFind(m, r, t)
			} else {
				r = m - 1
				continue
			}
		}
	}

	return false
}

func Search2() {
	fmt.Printf("<081> ")
	fmt.Println(search2([]int{1, 3, 1, 1, 1}, 3))
}

