package problems

import "fmt"

func wordBreak2(s string, wordDict []string) []string {
	d := []byte(s)
	w := wordDict

	if len(s) == 0 || len(w) == 0 {
		return nil
	}

	smap := make(map[string]int)

	for _, v := range w {
		smap[v] = 1
	}

	buf := make([]string, 0)
	out := make([]string, 0)

	var wordBreakDfs func (b []byte, buf *[]string)
	wordBreakDfs = func (b []byte, buf *[]string) {
		if len(b) == 0 {
			ps := ""

			for i, v := range *buf {
				if i == 0 {
					ps = ps + v
				} else {
					ps = ps + " " + v
				}
			}

			out = append(out, ps)
			return
		}

		for k := len(b) - 1; k >= 0; k-- {
			cad := string(b[:k + 1])
			_, ok := smap[cad]

			if ok {
				*buf = append(*buf, cad)
				wordBreakDfs(b[k + 1:], buf)
				*buf = (*buf)[:len(*buf) - 1]
			}
		}
	}

	wordBreakDfs(d, &buf)

	return out
}

func WordBreak2() {
	s := "catsanddog"
	w := []string {"cat", "cats", "and", "sand", "dog"}

	fmt.Printf("<140> will TLE ")
	fmt.Printf("%q\n", wordBreak2(s, w))
}

