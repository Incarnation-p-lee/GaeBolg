package problems

import "fmt"
import "math"

func candy(ratings []int) int {
	r := ratings

	if len(r) == 0 {
		return 0
	}

	r = append([]int{math.MaxInt32}, r...)
	r = append(r, math.MaxInt32)
	cost := make([]int, len(r))
	queue := make([]int, 0)

	var maxInt func (a, b int) int
	maxInt = func (a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	for i := 1; i < len(r) - 1; i++ {
		if r[i] <= r[i + 1] && r[i] <= r[i - 1] {
			queue = append(queue, i)
			cost[i] = 1

			for k := i + 1; k < len(r) - 1; k++ {
				if r[k] != r[i] {
					if k > i + 1 {
						queue = append(queue, k - 1)
						i = k - 1
					}
					break
				}

				cost[k] = 1
			}
		}
	}

	for _, v := range queue {
		for i := v; i > 0; i-- {
			if r[i - 1] < r[i] {
				break
			} else if r[i - 1] > r[i] {
				cost[i - 1] = maxInt(cost[i] + 1, cost[i - 1])
			}
		}

		for i := v; i < len(r) - 1; i++ {
			if r[i + 1] < r[i] {
				break
			} else if r[i + 1] > r[i] {
				cost[i + 1] = maxInt(cost[i] + 1, cost[i + 1])
			}
		}
	}

	total := 0

	for i := 1; i < len(r) - 1; i++ {
		total += cost[i]
	}

	return total
}

func Candy() {
	data := []int {4, 1, 2, 3, 4, 3, 5, 2, 4}

	fmt.Printf("<135> ")
	fmt.Println(candy(data))
}

