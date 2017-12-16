package problems

import "fmt"

func rotateArray(nums []int, k int) {
	n := nums
	k = k % len(nums)

	if len(nums) == 0 || k == 0 {
		return
	}

	var reverse func (l, r int)
	reverse = func (l, r int) {
		for l < r {
			n[l], n[r] = n[r], n[l]

			l++
			r--
		}
	}

	reverse(0, len(n) - k - 1)
	reverse(len(n) - k, len(n) - 1)
	reverse(0, len(n) - 1)
}

func RotateArray() {
	data := []int {1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("<189> ")
	rotateArray(data, 3)
	fmt.Println(data)
}

