package problems

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var length int
	data := []byte(s)

	if (s == "") {
		return 0
	}

	length = 1
	data = append(data, data[len(data) - 1])

	for i := 0; i < len(data); i++ {
		smap := make(map[byte]int)

		for j := i; j < len(data); j++ {
			d := data[j]

			if _, ok := smap[d]; !ok {
				smap[d] = j
			} else {
				if j - i > length {
					length = j - i
				}

				i = smap[d]
				break
			}
		}
	}

	return length
}

func lengthOfLongestSubstring2(s string) int {
	data := []byte(s)
	var length, left int
	var index [256]int

	left = -1

	for i, _ := range index {
		index[i] = -1
	}

	for i, v := range data {
		if index[v] != -1 {
			left = index[v]
		}

		if i - left > length {
			length = i - left
		}

		index[v] = i
	}

	return length
}

func LengthOfLongestSubstring() {
	var s string = "ohvhjdml"

	fmt.Printf("<003> ")
	fmt.Println(lengthOfLongestSubstring(s))
	fmt.Printf("<003> ")
	fmt.Println(lengthOfLongestSubstring2(s))
}

