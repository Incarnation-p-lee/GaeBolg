package problems

import "fmt"

func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	n1, n2 := nums1, nums2

	if m == 0 || len(n1) == 0 {
		return
	} else if n == 0 || len(n2) == 0 {
		return
	}

	d := make([]int, m)
	copy(d, n1)

	k := 0
	i, j := 0, 0

	for i < m && j < n {
		if d[i] < n2[j] {
			n1[k] = d[i]
			k++
			i++
		} else {
			n1[k] = n2[j]
			k++
			j++
		}
	}

	for i < m {
		n1[k] = d[i]
		k++
		i++
	}

	for j < n {
		n1[k] = n2[j]
		k++
		j++
	}
}

func MergeSortedArray() {
	n1 := []int {2, 6, 8, 9, 0, 0, 0}
	n2 := []int {1, 3, 7}

	fmt.Printf("<088> ")
	mergeSortedArray(n1, 4, n2, 3)
	fmt.Println(n1)
}

