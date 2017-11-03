package problems

import "fmt"
import "sort"

func isTransform(s1, s2 string) bool {
	d1 := []byte(s1)
	d2 := []byte(s2)
	count := 0

	for i, v := range d1 {
		if d2[i] != v {
			count++
		}
	}

	if count == 1 {
		return true
	} else {
		return false
	}
}

func findLaddersDfs(dirt []string, i int, deep int, b, e string, buf *[]string,
	out *[][]string, visited *[]bool, pMap *map[string]int) {
	if i == deep {
		return
	}

	var foundAdd func (buf *[]string)
	foundAdd =  func (buf *[]string) {
		m := make([]string, len(*buf))
		copy(m, *buf)
		*out = append(*out, m)
	}

	if b == e {
		foundAdd(buf)
		return
	}

	for j := 0; j < len(dirt); j++ {
		d := dirt[j]
		if _, ok := (*pMap)[b + d]; ok && !(*visited)[j] {
			*buf = append(*buf, d)
			(*visited)[j] = true
			findLaddersDfs(dirt, i + 1, deep, d, e, buf, out, visited, pMap)
			*buf = (*buf)[:len(*buf) - 1]
			(*visited)[j] = false
		}
	}
}

func findLaddersBfs(data []string, b, e string, pMap *map[string]int) int {
	var queue, slave []string
	var imap map[string]int
	var visited []bool
	var is_last bool

	val, deep := 1, 1
	imap = make(map[string]int)
	visited = make([]bool, len(data))
	queue = append(queue, b)

	for i, v := range data {
		imap[v] = i
	}

	for {
		if len(queue) == 0 {
			if is_last {
				break
			}

			deep++

			if len(slave) == 0 {
				break
			}

			queue, slave = slave, queue
		}

		d := queue[0]
		queue = queue[1:]

		if d == e {
			is_last = true
		}

		if d != b {
			visited[imap[d]] = true
		}

		for i, v := range data {
			if isTransform(d, v) && !visited[i] {
				(*pMap)[d + v] = val
				val++
				slave = append(slave, v)
			}
		}
	}

	return deep
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	var out [][]string
	var buf []string
	var visited []bool
	var deep int
	var pMap map[string]int

	sort.Strings(wordList)

	pMap = make(map[string]int)
	deep = findLaddersBfs(wordList, beginWord, endWord, &pMap)

	buf = append(buf, beginWord)
	visited = make([]bool, len(wordList))

	findLaddersDfs(wordList, 0, deep, beginWord, endWord, &buf, &out, &visited, &pMap)

	return out
}

func FindLadders() {
	b, e := "hit", "cog"
	dirt := []string{"hot","cog","dog","lot","log","dot"}

	fmt.Printf("<126> Will Time Limit Exceed ")
	fmt.Println(findLadders(b, e, dirt))
}

