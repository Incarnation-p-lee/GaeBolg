package problems

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	} else if len(matrix) == 1 {
		return matrix[0]
	}

	out := make([]int, 0)

	if len(matrix[0]) == 1 {
		for _, v := range matrix {
			out = append(out, v[0])
		}

		return out
	}

	n, m := len(matrix), len(matrix[0])
	mmap := make(map[string]int)

	i, j := 0, 0
	for ; j < m; j++ {
		out = append(out, matrix[i][j])
		mmap[string([]byte{byte(i), '-', byte(j)})] = j
	}

	i, j = 1, len(matrix[0]) - 1
	for ; i < n; i++ {
		out = append(out, matrix[i][j])
		mmap[string([]byte{byte(i), '-', byte(j)})] = j
	}

	i, j = len(matrix) - 1, len(matrix[0]) - 2
	for ; j >= 0; j-- {
		out = append(out, matrix[i][j])
		mmap[string([]byte{byte(i), '-', byte(j)})] = j
	}

	i, j = len(matrix) - 2, 0
	for ; i > 0; i-- {
		out = append(out, matrix[i][j])
		mmap[string([]byte{byte(i), '-', byte(j)})] = j
	}

	if n < 3 || m < 3 {
		return out
	}

	/*
	 * 0 - right
	 * 1 - down
	 * 2 - left
	 * 3 - up
	 */
	direct := 0
	i, j = 1, 1

	for {
		s := string([]byte{byte(i), '-', byte(j)})

		if _, ok := mmap[s]; ok {
			break
		}

		out = append(out, matrix[i][j])
		mmap[s] = i

		switch direct {
		case 0:
			s := string([]byte{byte(i), '-', byte(j + 1)})
			if _, ok := mmap[s]; !ok {
				j++
			} else {
				direct = 1
				i++
			}
		case 1:
			s := string([]byte{byte(i + 1), '-', byte(j)})
			if _, ok := mmap[s]; !ok {
				i++
			} else {
				direct = 2
				j--
			}
		case 2:
			s := string([]byte{byte(i), '-', byte(j - 1)})
			if _, ok := mmap[s]; !ok {
				j--
			} else {
				direct = 3
				i--
			}
		case 3:
			s := string([]byte{byte(i - 1), '-', byte(j)})
			if _, ok := mmap[s]; !ok {
				i--
			} else {
				direct = 0
				j++
			}
		}
	}

	return out
}

func SpiralOrder() {
	data := [][]int {
		[]int{ 1,  2,  3},
		[]int{ 8,  9,  4},
		[]int{ 7,  6,  5},
		[]int{ 0, 13, 14},
		[]int{10, 11, 12},
	}

	fmt.Printf("<054> ")
	fmt.Println(spiralOrder(data))
}

