package problems

import "fmt"

type Line struct {
	A, B, C int
}

/* Ax + B = Cy */

func maxPoints(points []Point) int {
	p := points

	if len(p) == 0 {
		return 0
	} else if len(p) == 1 {
		return 1
	}

	max := 0
	smap := make(map[Line]int)
	data := make([]Line, 0)

	for i, v1 := range p {
		for j, v2 := range p {
			if i == j || (v1.X == v2.X && v1.Y == v2.Y) {
				continue
			}

			a, b, c := 0, 0, 0

			if v1.X == v2.X {
				a, b, c = -1, v1.X, 0
			} else if v1.Y == v2.Y {
				a, b, c = 0, v1.Y, 1
			} else {
				a, b, c = v2.Y - v1.Y, v1.Y * v2.X - v2.Y * v1.X, v2.X - v1.X
			}

			s := Line{a, b, c}
			_, ok := smap[s]

			if !ok {
				smap[s] = 1
				data = append(data, s)
			}
		}
	}

	if len(data) == 0 {
		return len(p)
	}

	cad := make([]int, len(data))

	for _, v := range p {
		for i, s := range data {
			if s.A * v.X + s.B == s.C * v.Y {
				cad[i]++
			}
		}
	}

	for _, v := range cad {
		if v > max {
			max = v
		}
	}

	return max
}

func MaxPoints() {
	data := []Point {
		{1, 1}, {1, 2}, {1, 1}, {2, 2}, {3, 3}, {3, 0}, {2, 4},
	}

	fmt.Printf("<149> ")
	fmt.Println(maxPoints(data))
}

