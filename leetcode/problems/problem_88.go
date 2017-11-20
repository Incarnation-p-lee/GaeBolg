package problems

import "fmt"

func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	n1, n2 := nums1, nums2
	k := m + n - 1
	i, j := m - 1, n - 1

	for i >= 0 && j >= 0 {
		if n1[i] < n2[j] {
			n1[k] = n2[j]
			j--
		} else {
			n1[k] = n1[i]
			i--
		}

		k--
	}

	for i >= 0 {
		n1[k] = n1[i]
		k--
		i--
	}

	for j >= 0 {
		n1[k] = n2[j]
		k--
		j--
	}
}

func MergeSortedArray() {
	n1 := []int {2, 6, 8, 9, 0, 0, 0}
	n2 := []int {1, 3, 7}

	fmt.Printf("<088> ")
	mergeSortedArray(n1, 4, n2, 3)
	fmt.Println(n1)
}

