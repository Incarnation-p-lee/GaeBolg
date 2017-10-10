package main

import "os"
import "fmt"
import "strings"
import "bufio"
import "io/ioutil"
import "image"
import "image/color"
import "image/gif"
import "math"
import "math/rand"
import "net/http"

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
 * 12. ReadFile return byte slice(something or array?), need to convert to string.
 * 13. Bufi.Scanner and ioutil.ReadFile are implemented by os.File read and write.
 * 14. []color.Color{} for slice and gif.GIF{} for structure, append internal function.
 *
 */

var is_print bool = false

func main() {
	fmt.Println("Hello, World!")
	os_args()
	os_args_range()
	os_args_join()
	os_args_raw()

	bufio_uniq()
	bufio_file()
	ioutil_readfile()
	image_gif()
	net_http_url_request()
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
			/* bufio.NewScanner(fop) */
			fop.Close()
		}
	}
}

func ioutil_readfile() {
	counts := make(map[string]int)
	for _, filename := range os.Args[0:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail to open %s:%d\n", filename, err)
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
}

func image_gif() {
	var palette = []color.Color{color.White, color.Black}

	const (
		whiteIndex = 0 /* first color in palette */
		blackIndex = 1
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles * 2 * math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(size + int(x * size + 0.5),
				size + int(y * size + 0.5),
				blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	if is_print {
		gif.EncodeAll(os.Stdout, &anim)
	}
}

func net_http_url_request() {
	url := "http://www.bing.com"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fetch %s:%d fail\n", url, err)
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Read fail %d\n", err)
		os.Exit(1)
	}

	if is_print {
		fmt.Printf("%s", b)
	}
}

