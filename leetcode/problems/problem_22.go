package problems

import "fmt"

/* Generate Parentheses */

func GenerateParenthesis() {
	fmt.Printf("<022> ")
	fmt.Println(generateParenthesis(4))
}

/*
 * l means used left  (
 * r means used right )
 */
func generateDfs(l, r, n int, s string, ret *[]string) {
	if l == n && r == n {
		*ret = append(*ret, s)
	}

	if l < n { /* left ( has rest */
		generateDfs(l + 1, r, n, s + "(", ret)
	}

	if l > r { /* left > right, so can add right */
		generateDfs(l, r + 1, n, s + ")", ret)
	}
}

/*
 * Consider to build the 2*n depth binary tree, and filter some of the paths.
 */
func generateParenthesis(n int) []string {
	var out []string

	if n < 1 {
		return out
	}

	generateDfs(0, 0, n, "", &out)

	return out;
}

