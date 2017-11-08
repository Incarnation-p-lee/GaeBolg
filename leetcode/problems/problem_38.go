package problems

import "fmt"
import "strconv"

func countAndSay(n int) string {
	var data, slave []byte
	if n < 1 {
		return ""
	} else if n == 1 {
		return "1"
	} else if n == 2 {
		return "11"
	}

	data = append(data, '1', '1')

	var countAppend func (count int, v byte)
	countAppend = func (count int, v byte) {
		s := strconv.Itoa(count)
		slave = append(slave, []byte(s)...)
		slave = append(slave, v)
	}

	count := 1

	for i := 3; i <= n; i++ {
		for j, v := range(data) {
			if j > 0 && v == data[j - 1] {
				count++
			}

			if j > 0 && v != data[j - 1] {
				countAppend(count, data[j - 1])
				count = 1
			}

			if j == len(data) - 1 {
				countAppend(count, data[j])
				count = 1
			}
		}

		data = make([]byte, len(slave))
		copy(data, slave)
		slave = slave[:0]
	}

	return string(data)
}

func CountAndSay() {
	fmt.Printf("<038> ")
	fmt.Println(countAndSay(5))
}

