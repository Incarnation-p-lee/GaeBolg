package problems

import "fmt"
import "sort"
import "strconv"

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}

	var numLength func (a int) int
	numLength = func (a int) int {
		out := 1

		for a >= 10 {
			out++
			a = a / 10
		}

		return out
	}

	var appendNum func (a, b int) int
	appendNum = func (a, b int) int {
		l := numLength(b)

		for l != 0 {
			a = a * 10
			l--
		}

		return a + b
	}

	n := make([]int, len(nums))
	copy(n, nums)

	sort.Slice(n, func (i, j int) bool {
		a, b := n[i], n[j]
		d1, d2 := appendNum(a, b), appendNum(b, a)

		if d1 > d2 {
			return true
		} else {
			return false
		}
	});

	out := ""

	for _, v := range n {
		if out == "0" {
			out = ""
		}
		out += strconv.Itoa(v)
	}

	return out
}

func LargestNumber() {
	fmt.Printf("<179> ")
	fmt.Println(largestNumber([]int {0, 0, 0, 0, 0}))
}

