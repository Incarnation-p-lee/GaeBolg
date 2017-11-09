package problems

import "fmt"
import "sort"

func groupAnagrams(strs []string) [][]string {
	var out [][]string

	var computeMask func (s string) [26]int
	computeMask = func (s string) [26]int {
		var mask [26]int

		for _, v := range []byte(s) {
			mask[int(v - 'a')]++
		}

		return mask
	}

	smap := make(map[[26]int]int)

	for _, s := range strs {
		mask := computeMask(s)
		i, ok := smap[mask]

		if !ok {
			i = len(out)
			smap[mask] = i
			out = append(out, []string{})
		}

		out[i] = append(out[i], s)
	}

	for _, v := range out {
		sort.Strings(v)
	}

	sort.Slice(out, func (a, b int) bool {
		if len(out[a]) < len(out[b]) {
			return true
		} else {
			return false
		}
	})

	return out
}

func GroupAnagrams() {
	fmt.Printf("<049> ")
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

