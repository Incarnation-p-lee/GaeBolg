package problems

import "fmt"
import "strconv"

func restoreIpAddressesDfs(data []byte, idx, deep int, buf *[]byte, out *[]string) {
	if idx == len(data) && deep == 4 {
		s := string((*buf)[:len(*buf) - 1])
		*out = append(*out, s)
	} else if (idx < len(data) && deep < 4) {
		d := data[idx]

		*buf = append(*buf, d, byte('.'))
		restoreIpAddressesDfs(data, idx + 1, deep + 1, buf, out)
		*buf = (*buf)[:len(*buf) - 2]

		if idx + 1 < len(data) {
			if v, _ := strconv.Atoi(string(data[idx:idx + 2])); v > 9 {
				*buf = append(*buf, d, data[idx + 1], '.')
				restoreIpAddressesDfs(data, idx + 2, deep + 1, buf, out)
				*buf = (*buf)[:len(*buf) - 3]
			}
		}

		if idx + 2 < len(data) {
			if v, _ := strconv.Atoi(string(data[idx:idx + 3])); v < 256 && v > 99 {
				*buf = append(*buf, d, data[idx + 1], data[idx + 2], '.')
				restoreIpAddressesDfs(data, idx + 3, deep + 1, buf, out)
				*buf = (*buf)[:len(*buf) - 4]
			}
		}
	}
}

func restoreIpAddresses(s string) []string {
	var out []string
	var buf []byte

	restoreIpAddressesDfs([]byte(s), 0, 0, &buf, &out)
 
	return out
}

func RestoreIpAddresses() {
	s := "102550255"

	fmt.Printf("<093> ")
	fmt.Println(restoreIpAddresses(s)) 
}

