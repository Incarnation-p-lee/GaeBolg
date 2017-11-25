package problems

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	b, e, w := beginWord, endWord, wordList

	if len(b) == 0 || len(e) == 0 {
		return 0
	}

	isExist := false

	for _, s := range w {
		if e == s {
			isExist = true
		}
	}

	if !isExist {
		return 0
	}

	var canTrans func (a, b string) bool
	canTrans = func (a, b string) bool {
		count := 0
		da, db := []byte(a), []byte(b)

		for i, va := range da {
			if va != db[i] {
				count++
			}
		}

		if count == 1 {
			return true
		} else {
			return false
		}
	}

	var processNext func (v []int, i int, s, n string, m map[string]int, q *[]string)
	processNext = func (v []int, i int, s, n string, m map[string]int, q *[]string) {
		if v[i] == 0 {
			if canTrans(s, n) {
				*q = append(*q, n)
				m[n] = 1
				v[i] = 1
			}
		}
	}

	q1, q2 := make([]string, 0), make([]string, 0)
	s1, s2 := make([]string, 0), make([]string, 0)
	map1, map2 := make(map[string]int), make(map[string]int)
	v1, v2 := make([]int, len(w)), make([]int, len(w))

	q1 = append(q1, b)
	q2 = append(q2, e)
	deep1, deep2 := 0, 0

	for len(q1) != 0 && len(q2) != 0 {
		for len(q1) != 0 {
			d1 := q1[0]
			q1 = q1[1:]
			delete(map1, d1)

			if _, ok := map2[d1]; ok {
				return deep1 + deep2 + 1
			}

			for i, v := range w {
				processNext(v1, i, d1, v, map1, &s1)
			}
		}

		deep1++
		q1, s1 = s1, q1

		for len(q2) != 0 {
			d2 := q2[0]
			q2 = q2[1:]
			delete(map2, d2)

			if _, ok := map1[d2]; ok {
				return deep1 + deep2 + 1
			}

			for i, v := range w {
				processNext(v2, i, d2, v, map2, &s2)
			}
		}

		deep2++
		q2, s2 = s2, q2
	}

	return 0
}

func LadderLength() {
	b, e := "hit", "cog"
	wlist := []string{"hot", "dot", "dog", "lot", "log", "cog", "cot"}

	fmt.Printf("<127> ")
	fmt.Println(ladderLength(b, e, wlist))
}

