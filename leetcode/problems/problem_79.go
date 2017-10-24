package problems

import "fmt"

func isBoardVisited(d [][]byte, i, j int) bool {
	v := uint8(d[i][j])

	if v >> 7 == 1 {
		return true
	} else {
		return false
	}
}

func boardSetVisited(d [][]byte, i, j int) {
	v := uint8(d[i][j])
	v = v + (1 << 7)

	d[i][j] = byte(v)
}

func boardSetUnvisited(d [][]byte, i, j int) {
	v := uint8(d[i][j])
	v = v & 0x7f

	d[i][j] = byte(v)
}

func existDfs(d [][]byte, t []byte, k int, i, j int) bool {
	if (k == len(t)) {
		return true
	} else if i < 0 || j < 0 || i == len(d) || j == len(d[0]) {
		return false
	} else if isBoardVisited(d, i, j) {
		return false
	} else if (t[k] != d[i][j]) {
		return false
	} else {
		boardSetVisited(d, i, j)

		if found := existDfs(d, t, k + 1, i + 1, j,   ); found {
			return true
		}


		if found := existDfs(d, t, k + 1, i,     j + 1); found {
			return true
		}

		if found := existDfs(d, t, k + 1, i - 1, j,   ); found {
			return true
		}

		if found := existDfs(d, t, k + 1, i,     j - 1); found {
			return true
		}

		boardSetUnvisited(d, i, j)

		return false
    }
}

func exist(board [][]byte, word string) bool {
	bword := []byte(word) 
	var data [][]byte

	for _, row := range board {
		v := make([]byte, len(row))
		data = append(data, v)
	}

	for i, row := range board {
		for j, _ := range row {
			if board[i][j] != bword[0] {
				continue
			}

			for i, row := range board {
				copy(data[i], row)
			}

			if is_exist := existDfs(data, bword, 0, i, j); is_exist {
				return true
			}
		}
	}

	return false
}

func Exist() {
	board := [][]byte {
		{'A','B','C','E'},
		{'S','F','E','S'},
		{'A','D','E','E'},
	}

	fmt.Printf("<79> ")
	fmt.Println(exist(board, "ABCESEEEFS"))
}

