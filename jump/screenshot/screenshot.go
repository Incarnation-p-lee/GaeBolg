package screenshot

import "os"
import "fmt"
import "image"
import "image/color"
import _ "image/png"
import "bufio"
import "math"

var targetStartY int = 384
var topToCenter int = 188
var topToLeft int = 31
var matrixBuf [][]uint32 = make([][]uint32, 0)

func createRGBAArray(width, height int) [][]uint32 {
	if width <= 0 || height <= 0 {
		return nil
	}

	if len(matrixBuf) != 0 {
		return matrixBuf
	}

	w, h := width, height
	matrixBuf = make([][]uint32, h)

	for i, _ := range matrixBuf {
		matrixBuf[i] = make([]uint32, w)
	}

	return matrixBuf
}

func imageToIntMatrix(path string) [][]uint32 {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Open image %s done.\n", path)

	img, format, err := image.Decode(bufio.NewReader(file))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var fillRGBAArray func(matrix [][]uint32)
	fillRGBAArray = func(matrix [][]uint32) {
		rgba := img.(*image.NRGBA)

		for i, row := range matrix { /* Y height */
			for j, _ := range row { /* X width */
				col := rgba.At(j, i).(color.NRGBA)
				r, g, b := uint32(col.R), uint32(col.G), uint32(col.B)
				matrix[i][j] = uint32((r << 16) + (g << 8) + b)
			}
		}
	}

	rect := img.Bounds()
	fmt.Printf("Decode %s [%d * %d] done.\n", format, rect.Max.X, rect.Max.Y)

	out := createRGBAArray(rect.Max.X, rect.Max.Y)
	fillRGBAArray(out)

	return out
}

func absUint32(a, b uint32) uint32 {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func absInt(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func isSameRegion(m, n uint32) bool {
	var maxdiff uint32 = 25
	rm, gm, bm := m>>16, (m>>8)&0xff, m&0xff
	rn, gn, bn := n>>16, (n>>8)&0xff, n&0xff

	if absUint32(rm, rn) > maxdiff {
		return false
	} else if absUint32(gm, gn) > maxdiff {
		return false
	} else if absUint32(bm, bn) > maxdiff {
		return false
	} else {
		return true
	}
}

func findTargetPoint(matrix [][]uint32, ref uint32, limit int) (int, int) {
	x1, y1, y2, m := 0, 0, 0, matrix

	if len(m) == 0 || len(m[0]) == 0 {
		return 0, 0
	}

Loop1:
	for i := targetStartY; i < len(m); i++ {
		row := m[i]
		if limit > (len(m[0]) / 2) { /* Ball on right */
			for j := 0; j < (limit - topToLeft); j++ {
				if isSameRegion(row[j], ref) == false {
					x1, y1 = j, i
					break Loop1
				}
			}
		} else { /* Ball on left */
			for j := limit + topToLeft; j < len(m[0]); j++ {
				if isSameRegion(row[j], ref) == false {
					x1, y1 = j, i
					break Loop1
				}
			}
		}

	}

	k := x1
	for isSameRegion(m[y1][k], ref) && k < x1+8 {
		k++
	}

	widthMax, l, r := 0, x1, k

	for i := y1; i < len(m); i++ {
		isLeftSame, isRightSame := false, false
		for l >= 0 && r < len(m[0]) {
			isLeftSame = isSameRegion(m[i][l], ref)
			isRightSame = isSameRegion(m[i][r], ref)

			if isLeftSame && isRightSame {
				break
			}
			if !isLeftSame {
				l--
			}
			if !isRightSame {
				r++
			}
		}

		if (r - l) == widthMax {
			y2 = i
			break
		}
		widthMax = r - l
	}

	fmt.Printf("Target [%d, %d]\n", x1, y2)
	return x1, y2
}

func isBlackMain(v uint32) bool {
	r, g, b := v>>16, (v>>8)&0xff, v&0xff

	if r != 0x34 {
		return false
	} else if g != 0x35 {
		return false
	} else if b != 0x3b {
		return false
	} else {
		return true
	}
}

func findSourcePoint(matrix [][]uint32, ref uint32) (int, int) {
	m := matrix

	if len(m) == 0 || len(m[0]) == 0 {
		return 0, 0
	}

	x1, y1 := 0, 0

Loop1:
	for i := targetStartY; i < len(m); i++ {
		row := m[i]
		for j, v := range row {
			if isBlackMain(v) {
				x1, y1 = j+5, i
				break Loop1
			}
		}
	}

	y1 += topToCenter

	fmt.Printf("Source [%d, %d]\n", x1, y1)

	return x1, y1
}

func computeDistance(matrix [][]uint32) int {
	m := matrix

	if len(m) == 0 || len(m[0]) == 0 {
		return 0
	}

	background := func() uint32 {
		var last uint32

		for i := 0; i < 5; i++ {
			row := m[i]
			for _, v := range row {
				if last == 0 {
					last = v
				} else if v != last {
					fmt.Println("Mismatched background!")
				}
			}
		}

		return last
	}()

	x1, y1 := findSourcePoint(matrix, background)
	x2, y2 := findTargetPoint(matrix, background, x1)

	a := float64(absInt(x1, x2))
	b := float64(absInt(y1, y2))

	distance := math.Sqrt(a*a + b*b)
	fmt.Printf("Computed distance %d\n", int(distance))

	return int(distance)
}

func Distance(path string) int {
	if path == "" {
		return -1
	}

	matrix := imageToIntMatrix(path)
	distance := computeDistance(matrix)

	return distance
}
