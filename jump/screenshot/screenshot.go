package screenshot

import "os"
import "fmt"
import "image"
import "image/color"
import _ "image/png"
import "bufio"
import "math"

var targetStartRatio float64 = 0.28654
var targetRatio float64 = 0.09792
var targetLeftRatio float64 = 0.01823
var targetColor uint32 = 0x34353b
var gameOverBackground uint32 = 0x332e2c
var topMax int = 16

var matrixBuf [][]uint32 = make([][]uint32, 0)

func targetStartRow(height int) int {
	return int(float64(height) * targetStartRatio)
}

func targetCenterOffset(height int) int {
	return int(float64(height) * targetRatio)
}

func targetLeftToCenter(height int) int {
	return int(float64(height) * targetLeftRatio)
}

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
	out := createRGBAArray(rect.Max.X, rect.Max.Y)
	fillRGBAArray(out)

	fmt.Printf("Decode %s [%d * %d] done.\n", format, rect.Max.X, rect.Max.Y)

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
	var maxDiff uint32 = 20
	rm, gm, bm := (m >> 16), ((m >> 8) & 0xff), (m & 0xff)
	rn, gn, bn := (n >> 16), ((n >> 8) & 0xff), (n & 0xff)

	if absUint32(rm, rn) > maxDiff {
		return false
	} else if absUint32(gm, gn) > maxDiff {
		return false
	} else if absUint32(bm, bn) > maxDiff {
		return false
	} else {
		return true
	}
}

func findTargetPoint(matrix [][]uint32, ref uint32, limit int) (int, int) {
	x1, y1, y2, k, m := 0, 0, 0, 0, matrix

Loop1:
	for i := targetStartRow(len(m)); i < len(m); i++ {
		row := m[i]
		if limit > (len(m[0]) / 2) { /* Ball on right */
			for j := 0; j < (limit - targetLeftToCenter(len(m))); j++ {
				if isSameRegion(row[j], ref) == false {
					x1, y1 = j, i
					break Loop1
				}
			}
		} else { /* Ball on left */
			for j := limit + targetLeftToCenter(len(m)); j < len(m[0]); j++ {
				if isSameRegion(row[j], ref) == false {
					x1, y1 = j, i
					break Loop1
				}
			}
		}
	}

	for k = x1; isSameRegion(m[y1][k], m[y1][x1]) && k < (topMax+x1); k++ {
	}

	last0, last1, r, x1 := 0, 0, ((x1 + k) / 2), ((x1 + k) / 2)

	for i := y1; i < len(m); i++ {
		for ; r < len(m[0]); r++ {
			if isSameRegion(m[i][r], ref) {
				break
			}
		}

		if ((r - x1) == last0) && last0 == last1 {
			y2 = i
			break
		}

		last0 = last1
		last1 = r - x1
	}

	fmt.Printf("Target [%d, %d]\n", x1, y2)
	return x1, y2
}

func isBlackMain(v uint32) bool {
	maxDiff := uint32(1)
	r, g, b := (v >> 16), ((v >> 8) & 0xff), (v & 0xff)
	rt, gt, bt := (targetColor >> 16), ((targetColor >> 8) & 0xff),
		(targetColor & 0xff)

	if absUint32(r, rt) > maxDiff {
		return false
	} else if absUint32(g, gt) > maxDiff {
		return false
	} else if absUint32(b, bt) > maxDiff {
		return false
	} else {
		return true
	}
}

func findSourcePoint(matrix [][]uint32, ref uint32) (int, int) {
	m := matrix

	x1, y1 := 0, 0

Loop1:
	for i := targetStartRow(len(m)); i < len(m); i++ {
		row := m[i]
		for j, v := range row {
			if isBlackMain(v) {
				x1, y1 = (j + 5), i
				break Loop1
			}
		}
	}

	y1 += targetCenterOffset(len(m))

	fmt.Printf("Source [%d, %d]\n", x1, y1)

	return x1, y1
}

func isGameOver(background uint32) bool {
	rgb, maxDiff := background, uint32(15)
	r, g, b := (rgb >> 16), ((rgb >> 8) & 0xff), (rgb & 0xff)

	if absUint32(r, 0x33) > maxDiff {
		return false
	} else if absUint32(g, 0x2d) > maxDiff {
		return false
	} else if absUint32(b, 0x26) > maxDiff {
		return false
	} else {
		return true
	}
}

func computeDistance(matrix [][]uint32) int {
	background := matrix[targetStartRow(len(matrix))][0]

	if isGameOver(background) {
		fmt.Println("Game over!!")
		os.Exit(0)
	}

	x1, y1 := findSourcePoint(matrix, background)
	x2, y2 := findTargetPoint(matrix, background, x1)

	a := float64(absInt(x1, x2))
	b := float64(absInt(y1, y2))

	distance := int(math.Sqrt((a * a) + (b * b)))

	fmt.Printf("Computed distance %d\n", distance)

	return distance
}

func Distance(path string) int {
	if path == "" {
		return -1
	}

	matrix := imageToIntMatrix(path)
	distance := computeDistance(matrix)

	return distance
}
