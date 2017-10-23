package problems

import "fmt"

/* Problem 17. Letter Combinations of a Phone Number */
func LetterCombination() {
	fmt.Printf("<17> ")
	letterCombination("239")
}

func letterCombination(digits string) []string {
	var combin, tmp []string
	data := []byte(digits)
	dmap := map[byte][]string {
		'2': {"a", "b", "c",},
		'3': {"d", "e", "f",},
		'4': {"g", "h", "i",},
		'5': {"j", "k", "l",},
		'6': {"m", "o", "n",},
		'7': {"p", "q", "r", "s",},
		'8': {"t", "u", "v",},
		'9': {"w", "x", "y", "z",},
	}

	for _, d := range data {
		if len(combin) == 0 {
			combin = dmap[d]
		} else {
			letters := dmap[d]
			for _, c := range combin {
				for _, l := range letters {
					tmp = append(tmp, c + l)
				}
			}
			combin = tmp
			tmp = nil
		}
	}

	fmt.Println(combin)

	return combin
}

