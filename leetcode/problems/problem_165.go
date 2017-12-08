package problems

import "fmt"
import "strings"
import "strconv"

func compareVersion(version1 string, version2 string) int {
	v1, v2 := version1, version2

	if v1 == "" && v2 == "" {
		return 0
	} else if v2 == "" {
		return 1
	} else if v1 == "" {
		return -1
	}

	s1, s2 := strings.Split(v1, "."), strings.Split(v2, ".")

	for i := 0; i < len(s1) && i < len(s2); i++ {
		d1, _ := strconv.Atoi(s1[i])
		d2, _ := strconv.Atoi(s2[i])

		if d1 > d2 {
			return 1
		} else if d1 < d2 {
			return -1
		}
	}

	var isTailZero func (s []string) bool
	isTailZero = func (s []string) bool {
		for _, v := range s {
			if d, _ := strconv.Atoi(v); d != 0 {
				return false
			}
		}

		return true
	}

	if len(s1) < len(s2) {
		if isTailZero(s2[len(s1):]) {
			return 0
		} else {
			return -1
		}
	} else if len(s1) > len(s2) {
		if isTailZero(s1[len(s2):]) {
			return 0
		} else {
			return 1
		}
	} else {
		return 0
	}
}

func CompareVersion() {
	fmt.Printf("<165> ")
	fmt.Println(compareVersion("1.1", "1"))
}

