package problems

import "fmt"
import "math"

func minCut(s string) int {
	if len(s) == 0 {
		return 0
	}

	var isPalindrome func (d []byte, i, j int) bool
	isPalindrome = func (d []byte, i, j int) bool {
		if i > j {
			return false
		}

		l, r := i, j

		for l <= r {
			if d[l] != d[r] {
				return false
			}

			l++
			r--
		}

		return true
	}

	d := []byte(s)
	dpPalindrom := make([][]bool, len(d))

	for i, _ := range dpPalindrom {
		dpPalindrom[i] = make([]bool, len(d))
	}

	for i, _ := range d {
		for j, _ := range d {
			if i <= j {
				dpPalindrom[i][j] = isPalindrome(d, i, j)
			}
		}
	}

	dpCut := make([]int, len(d))
	dpCut[0] = 0

	for j := 1; j < len(d); j++ {
		min := math.MaxInt32

		if dpPalindrom[0][j] {
			min = 0
		} else {
			for i := 0; i < j; i++ {
				if dpPalindrom[i + 1][j] {
					count := dpCut[i] + 1
					if count < min {
						min = count
					}
				}
			}
		}

		dpCut[j] = min
	}

	return dpCut[len(d) - 1]
}

func MinCut() {
	fmt.Printf("<132> ")
	fmt.Println(minCut("abbaaacddc"))
}

