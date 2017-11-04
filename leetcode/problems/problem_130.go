package problems

import "fmt"

func solveBfs(data [][]byte, queue []int) {
	for len(queue) != 0 {
		k, l := queue[0], queue[1]
		data[k][l] = 'M'
		queue = queue[2:]

		if k > 0 && data[k - 1][l] == 'O' {
			queue = append(queue, k - 1, l)
		}

		if k < len(data) - 1 && data[k + 1][l] == 'O' {
			queue = append(queue, k + 1, l)
		}

		if l > 0 && data[k][l - 1] == 'O' {
			queue = append(queue, k, l - 1)
		}

		if l < len(data[0]) - 1 && data[k][l + 1] == 'O' {
			queue = append(queue, k, l + 1)
		}
	}
}

func solve(board [][]byte)  {
	var queue []int

	if len(board) < 2 || len(board[0]) < 2 {
		return
	}

	for k := 0; k < len(board[0]); k++ {
		if board[0][k] == 'O' {
			queue = append(queue, 0, k)
		}

		if board[len(board) - 1][k] == 'O' {
			queue = append(queue, len(board) - 1, k)
		}
	}

	for k := 1; k < len(board) - 1; k++ {
		if board[k][0] == 'O' {
			queue = append(queue, k, 0)
		}

		if board[k][len(board[0]) - 1] == 'O' {
			queue = append(queue, k, len(board[0]) - 1)
		}
	}

	solveBfs(board, queue)

	for i, d := range board {
		for j, v := range d {
			if v == 'O' {
				board[i][j] = 'X'
			} else if v == 'M' {
				board[i][j] = 'O'
			}
		}
	}
}

func Solve()  {
	var board [][]byte = [][]byte {
		[]byte{'X', 'O', 'X', 'O', 'X', 'O'},
		[]byte{'O', 'X', 'O', 'X', 'O', 'X'},
		[]byte{'X', 'O', 'X', 'O', 'X', 'O'},
		[]byte{'O', 'X', 'O', 'X', 'O', 'X'},
	}

	fmt.Printf("<130> ")
	solve(board)
	fmt.Println(board)
}

