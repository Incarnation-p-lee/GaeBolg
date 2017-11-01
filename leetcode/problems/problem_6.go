package problems

import "fmt"

func convert(s string, numRows int) string {
	var out string
	var slave [][]byte

	data := []byte(s)

	if len(s) == 0 {
		return s
	} else if numRows == 1 {
		return s
	} else if numRows == 2 {
		s1, s2 := "", ""
		for i, v := range data {
			if i % 2 == 0 {
				s1 += string(v)
			} else {
				s2 += string(v)
			}
		}

		return s1 + s2
	}

	for i := 0; i < numRows; i++ {
		var d []byte
		slave = append(slave, d)
	}

	j := 0
	up := false

	for _, v := range data {
		if !up {
			slave[j] = append(slave[j], v)
			j++

			if j == numRows {
				up = true
				j = numRows - 2
			}
		} else {
			slave[j] = append(slave[j], v)
			j--

			if j == 0 {
				up = false
			}
		}
	}

	for _, d := range slave {
		out += string(d)
	}

	return out
}

func Convert() {
	s := "PAYPALISHIRING"

	fmt.Printf("<006> ") 
	fmt.Println(convert(s, 3))
}

