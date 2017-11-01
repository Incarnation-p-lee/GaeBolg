package problems

import "fmt"

func reverse(x int) int {
	var n uint32
	var t, m uint64
	var d int64

	n = uint32(x)

	if x < 0 {
		n = ^n + 1
	}

	t = uint64(n)

	for ; t != 0; t = t / 10 {
		k := t % 10
		m = m * 10 + k

		if x > 0 && m > 2147483647 {
			return 0
		} else if x < 0 && m > 2147483648 {
			return 0
		}
	}

	d = int64(m)

	if x < 0 {
		d = -d
	}

	return int(d)
}

func Reverse() {
	fmt.Printf("<007> ") 
	fmt.Println(reverse(-123))
}


