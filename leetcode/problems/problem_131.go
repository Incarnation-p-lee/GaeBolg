package problems

import "fmt"

func partitionDfs(data []byte, i int, buf *[]string, out *[][]string) {
	if i == len(data) {
		m := make([]string, len(*buf))
		copy(m, *buf)
		*out = append(*out, m)
		return
	}

	var isPalindrome func (data []byte, s int, e int) bool
	isPalindrome = func (data []byte, s int, e int) bool {
		i, j := s, e
		for i < j {
			if data[i] != data[j] {
				return false
			} else {
				i++
				j--
			}
		}

		return true
	}

	for k := i; k < len(data); k++ {
		if isPalindrome(data, i, k) {
			*buf = append(*buf, string(data[i:k + 1]))
			partitionDfs(data, k + 1, buf, out)
			*buf = (*buf)[:len(*buf) - 1]
		}
	}
}

func partition(s string) [][]string {
	var buf []string
	var out [][]string

	partitionDfs([]byte(s), 0, &buf, &out)

	return out
}

func Partition() {
	fmt.Printf("<131> ")
	fmt.Println(partition("aabc"))
}

