package screenshot

import "os"
import "fmt"
import "image"
import "image/color"
import _ "image/png"
import "bufio"
import "math"

var targetStartY int = 384
var blockMaxY int = 300
var topToCenter int = 188

func createRGBAArray(width, height int) [][]uint32 {
	if width <= 0 || height <= 0 {
		return nil
	}

	w, h := width, height
	out := make([][]uint32, h)

	for i, _ := range out {
		out[i] = make([]uint32, w)
	}

	return out
}

func imageToIntAarry(path string) [][]uint32 {
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

	var fillRGBAArray func (matrix [][]uint32)
	fillRGBAArray = func (matrix [][]uint32) {
		rgba := img.(*image.NRGBA)

		for i, row := range matrix { /* Y height */
			for j, _ := range row {  /* X width */
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

func isSameRegion(m, n uint32) bool {
	var maxdiff uint32 = 20
	rm, gm, bm := m >> 16, (m >> 8) & 0xff, m & 0xff
	rn, gn, bn := n >> 16, (n >> 8) & 0xff, n & 0xff

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

func findTargetPoint(matrix [][]uint32, ref uint32) (int, int) {
	m := matrix

	if len(m) == 0 || len(m[0]) == 0 {
		return 0, 0
	}

	x1, y1, y2 := 0, 0, 0

Loop1:
	for i := targetStartY; i < len(m); i++ {
		row := m[i]
		for j, v := range row {
			if isSameRegion(v, ref) == false {
				x1, y1 = j, i
				break Loop1
			}
		}
	}

	if x1 == 0 && y1 == 0 {
		return 0, 0
	}

Loop2:
	for j := len(m[0]) - 1; j > x1; j-- {
		for i := y1; i < y1 + blockMaxY; i++ {
			if isSameRegion(m[i][j], ref) == false {
				y2 = i
				break Loop2
			}
		}
	}

	fmt.Printf("Target [%d, %d]\n", x1, y2)

	return x1, y2
}

func isBlackMain(v uint32) bool {
	r, g, b := v >> 16, (v >> 8) & 0xff, v & 0xff

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

func findSourcePoint(matrix [][]uint32, ref uint32, x, y int) (int, int) {
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
				x1, y1 = j + 5, i
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

	background := func () uint32 {
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
	} ()

	x1, y1 := findTargetPoint(matrix, background)
	x2, y2 := findSourcePoint(matrix, background, x1, y1)

	a := float64(absInt(x1, x2))
	b := float64(absInt(y1, y2))

	distance := math.Sqrt(a * a + b * b)
	fmt.Printf("Computed distance %d\n", int(distance))

	return int(distance)
}

func Distance(path string) int {
	if path == "" {
		return -1
	}

	matrix := imageToIntAarry(path)
	distance := computeDistance(matrix)

	return distance
}

