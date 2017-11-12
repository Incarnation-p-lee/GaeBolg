package problems

import "fmt"

func addBinary(a string, b string) string {
	if len(a) == 0 {
		return b
	} else if len(b) == 0 {
		return a
	}

	out := []byte{}
	isCarry := 0
	da, db := []byte(a), []byte(b)
	i, j := len(da) - 1, len(db) - 1

	for i >= 0 && j >= 0 {
		d := int(da[i] - '0') + int(db[j] - '0') + isCarry
		n := d % 2
		out = append([]byte{byte(n + '0')}, out...)

		if d > 1 {
			isCarry = 1
		} else {
			isCarry = 0
		}

		i--
		j--
	}

	for i >= 0 {
		d := int(da[i] - '0') + isCarry
		n := d % 2
		out = append([]byte{byte(n + '0')}, out...)

		if d > 1 {
			isCarry = 1
		} else {
			isCarry = 0
		}

		i--
	}

	for j >= 0 {
		d := int(db[j] - '0') + isCarry
		n := d % 2
		out = append([]byte{byte(n + '0')}, out...)

		if d > 1 {
			isCarry = 1
		} else {
			isCarry = 0
		}

		j--
	}

	if isCarry == 1 {
		out = append([]byte{'1'}, out...)
	}

	return string(out)
}

func AddBinary() {
	fmt.Printf("<067> ")
	fmt.Println(addBinary("11", "111101"))
}

