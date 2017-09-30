package main

import "os"
import "fmt"
import "strings"
import "bufio"
/* import ("fmt" "os") */

/*
 *  1. import must be located after package declaration.
 *  2. keyword func var const type.
 *  3. gofmt will convert source code to standard format.
 *  4. use tab for indent.
 *  5. static compiler.
 *  6. package os for interactive between different platform.
 *  7. slice s[m:n] equals s [m,n), s[m:] means until max size.
 *  8. automatically set uninitialized var to zero.
 *  9. ++/-- is statement not expression. Illegal if ++/--i.
 * 10. := only available when initialize, or variable without var.
 * 11. Almost all build-in function variable with first letter upper, os.Args.
 */

func main() {
	fmt.Println("Hello, World!")
	os_args()
	os_args_range()
	os_args_join()
	os_args_raw()

	bufio_uniq()
	bufio_file()
}

func os_args() {
	var s, sep string

	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}

func os_args_range() {
	s, sep := "", "" /* variable without var cannot be package variable */

        /*
         * FOREACH iteration, range will return two value of given array.
         * index and its value as a pair, (index, value).
         *
         * HERE use blank identifier _ for ignoring the returned index.
         */
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

func os_args_join() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

func os_args_raw() {
	fmt.Println(os.Args[0:])
}

func bufio_uniq() {
	/* map[key] = value */
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line);
		}
	}
}

func bufio_file() {
	files := os.Args[0:]

	for _, arg := range files {
		fop, err := os.Open(arg) /* return opened file and err */
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail to open %s:%d\n", arg, err)
		} else {
			fmt.Fprintf(os.Stdout, "Opened file %s\n", arg)
			fop.Close()
		}
	}
}

