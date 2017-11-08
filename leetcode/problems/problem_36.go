package problems

import "fmt"

func isValidSudoku(board [][]byte) bool {
	var rmap [9]map[byte]byte
	var bmap [9]map[byte]byte
	var cmap [9]map[byte]byte

	for i, _ := range bmap {
		bmap[i] = make(map[byte]byte)
		rmap[i] = make(map[byte]byte)
		cmap[i] = make(map[byte]byte)
	}

	for i, v := range board {
		for j, d := range v {
			if d != '.' {
				if _, ok := rmap[j][d]; ok {
					return false
				}

				rmap[j][d] = d

				if _, ok := cmap[i][d]; ok {
					return false
				}

				cmap[i][d] = d

				if _, ok := bmap[j / 3 + (i / 3) * 3][d]; ok {
					return false
				}

				bmap[j / 3 + (i / 3) * 3][d] = d
			}
		}
	}

	return true
}

func IsValidSudoku() {
	data := [][]byte {
		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},

		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},

		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Printf("<036> ")
	fmt.Println(isValidSudoku(data))
}

