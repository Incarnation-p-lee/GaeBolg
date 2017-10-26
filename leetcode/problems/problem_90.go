package problems

import "fmt"
import "sort"

func subsetsWithDupDfs(data []int, idx int, keep bool, buf *[]int, out *[][]int) {
	if idx == len(data) {
		m := make([]int, len(*buf))
		copy(m, *buf)
		*out = append(*out, m)
	} else {
		/* idx > 0 && data[idx] == data[idx - 1] && !keep, then skip */
		if idx == 0 || data[idx] != data[idx - 1] || keep {
			*buf = append(*buf, data[idx])
			subsetsWithDupDfs(data, idx + 1, true, buf, out)
			*buf = (*buf)[:len(*buf) - 1]
		}

		subsetsWithDupDfs(data, idx + 1, false, buf, out)
	}
}

func subsetsWithDup(nums []int) [][]int {
	var buf []int
	var out [][]int

	sort.Ints(nums)

	subsetsWithDupDfs(nums, 0, true, &buf, &out)

	sort.Slice(out, func (a, b int) bool {
			var sz int
			la := len(out[a])
			lb := len(out[b])

			if la > lb {
				sz = lb
			} else {
				sz = la
			}

			for i := 0; i < sz; i++ {
				if out[a][i] < out[b][i] {
					return true
				} else if out[a][i] > out[b][i] {
					return false
				}
			}

			if la < lb {
				return true
			} else {
				return false
			}
		})

	return out 
}

func SubsetsWithDup() {
	data := []int{2, 2, 2}

	fmt.Printf("<090> ")
	fmt.Println(subsetsWithDup(data))
}

