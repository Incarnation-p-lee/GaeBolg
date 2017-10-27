package problems

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var k, sz, v, last int 
	var need_next bool

	sz = len(nums1) + len(nums2)

	for i, j := 0, 0; true; k++ {
		if i == len(nums1) {
			v = nums2[j]
			j++
		} else if j == len(nums2) {
			v = nums1[i]
			i++
		} else if nums1[i] > nums2[j] {
			v = nums2[j]
			j++
		} else {
			v = nums1[i]
			i++
		}

		if need_next {
			return (float64(v) + float64(last)) / 2
		}

		if k == (sz - 1) / 2 {
			if (sz % 2) == 1 {
				return float64(v)
			} else {
				last = v
				need_next = true
			}
		}
	}

	return 0.0
}

func FindMedianSortedArrays() {
	var m []int = []int{1, 3, 4}
	var n []int = []int{2, 9, 10, 20}

	fmt.Printf("<004> ")
	fmt.Println(findMedianSortedArrays(m, n))
}

