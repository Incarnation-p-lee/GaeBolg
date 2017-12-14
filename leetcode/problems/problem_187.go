package problems

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	d := []byte(s)

	if len(d) == 0 {
		return []string{""}
	}

	out := make([]string, 0)
	smap := make(map[string]int)

	fmt.Println(len(d))

	for i := 0; i + 9 < len(d); i++ {
		cad := string(d[i:i + 10])

		if v, ok := smap[cad]; ok {
			if v == 1 {
				out = append(out, cad)
			}

			smap[cad] = v + 1
		} else {
			smap[cad] = 1
		}
	}

	return out
}

func FindRepeatedDnaSequences() {
	fmt.Printf("<187> ")
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
}

