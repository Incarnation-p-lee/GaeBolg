package main

import "os"
import "fmt"
import "strings"
import "bufio"
import "io"
import "io/ioutil"
import "image"
import "image/color"
import "image/gif"
import "math"
import "math/rand"
import "net/http"
import "time"
import "log"
import "sync"
import "flag"
import "bytes"
import "sort"

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
 * 15. goruntine is one concurrence execution way, and channel is for the commnucation
 *     between different goruntine. [go function] will create one new goruntime and then
 *     execution the function. when go function try to access channel in one function
 *     execution time, the channel need to handle the critical section itself.
 * 16. Pointer is available, but cannot do any operations on pointers.
 * 17. go doc http.ListenAndServe for local go manual.
 * 18. Symbol in package with beginning letter upper will be exported.
 * 19. No uninitialized var in go.
 * 20. Simple variable declartion := is available in function body.
 * 21. Agnist to = is assignments, := is for declarations.
 * 22. Point type *int, *float32.
 * 23. Return the pointer to local variable is valid.
 * 24. var p *int, *p++ will not change p, just *p = *p + 1
 * 25. Build-in function new(T) will create a var with type T, and return the pointer
 *     to this type. This is no diffence with declarations.
 * 26. Function args and return values are local variables.
 * 27. Function all like foo(a, b, c,) is valid.
 * 28. Go compiler will decide when to locate variable, heap or stack.
 * 29. Multi-variable assignment i, j = j, i is valid.
 * 30. Type cast T(x), only valid if they have the same low_level_type or pointers.
 * 31. Each package is one namespace, and each package contains many source files.
 *     Almost only one file contains the package comments(descriptions).
 * 32. One init function of package will be called when package initialize, this
 *     init function cannot be call and reference.
 * 33. Use simple declarations on the same name variable will cover the other variables.
 * 34. And Not ops, &^ (y AND (NOT x)), ^ on single ops means ~.
 * 35. Build-in function len return int type.
 * 36. && has higher priority to ||.
 * 37. String can be connected by +, and compared by == < >.
 * 38. String cannot be modified.
 * 39. String and byte slice can convert to each other.
 * 40. Build-in len cap real imag complex unsaft.Sizeof will compute at compiling time,
 *     and it will return const value.
 * 41. Const generator, itoa.
 * 42. Typeless const, if const has no clear type, go will provide at least 256-bits precision.
 * 43. If initialization with typeless const, there will be some implicit rules. Like typeless
 *     const will be treated as int, as well as float64 and complex128.
 * 44. For complicated data, there are array, slice, map, struct.
 * 45. Go will not convert array args to the pointer of the array when function calls.
 * 46. Slice is somewhat like an array without const lenght, []T, array is [L]T. Slice always
 *     contains 3 parts data, pointer, length and capacity.
 * 47. Passing slice to function as args is just pass a alias of slice, or copy one slice also
 *     just create one alias. Different to array, slice cannot be compared, only nil is valid.
 * 48. Build-in make able to create slice, make([]T, len, cap)
 * 49. Build-in append used to append new element to slice.
 * 50. Map is one reference to hash table, key should be compared.
 * 51. Struct member begin with upper letter will be exported.
 * 52. Struct can be initialized with signature, anim := git.GIF{LoopCountL: nf}, only exported
 *     member can use signature initialization.
 * 53. All function call will pass the copy of arguments, so pass the pointer of structure may
 *     be good for performance.
 * 54. If all member of struct is comparable, struct is comparable.
 * 55. Anonymous member of struct is only declared data type without a name. The language just
 *     give a default name of this member and ignore it when access. This always called struct
 *     embedded.
 * 56. Text/template can used to generate complicated text, as well as html/templete
 * 57. Function without body means it not implemented by Go.
 * 58. Go use dynamic stack for function call, no stack overflow.
 * 59. Function can return multi-values, and bare return, but not recommanded, as Go will
 *     initialize return value to zero when function begins.
 * 60. Build-in type error is a interface, and cannot be treated as exception.
 * 61. Function values, like other types, function also got values, or function pointer, which
 *     cannot be compared.
 * 62. Anonymous function is supported, for this anonymous function can access the outer
 *     function variables. Also function values are reference types. When anonymous function
 *     touch the local variable in caller function, it will save the address of that variable.
 * 63. Functions also support variadic parameters.
 * 64. Prefix defer when call functions. In one function with a statement defer, it will be 
 *     executed after all function finished. If function has more than one defer, it first
 *     one will be the last execution one. Defer always for resource open/close, connect/disconnect
 *     lock and unlock.
 * 65. Runtime error can use panic, will exit current goruntine and print log. Also panic can be
 *     used as assert, but should be fatal error. If deferred function invoked build-in function
 *     recover, recover will occurs from panic and return panic value, instead of exit.
 */

/*
 * Keyword
 *     break case   chan   const  continue default defer     else fallthrough
 *     for   func   go     goto   if       import  interface map  package
 *     range return select struct switch   type    var
 *
 * Pre-defined
 *     true    false   iota      nil
 *     int     int8    int16     int32      int64
 *     uint    uint8   uint16    uint32     uint64  uintptr
 *     float32 float64 complex64 complex128
 *     bool    byte    rune      string     error
 *     make    len     cap       new        append  copy    close delete
 *     complex real    imag      panic      recover
 *
 * Declaration
 *     var const type func
 *     func func_name(arg list) return list {
 *     }
 *     func (t T) func_name() string {
 *     }
 *     That indicate the func_name is for Type T, t.func_name.
 *     var var_name type = expression
 *
 * New type
 *     type type_name low_level_type
 *     type Celsius float64, now type Celsius is typedef to float64
 *     type Fahren  float64, now type Fahren  is typedef to float64
 *     NOTE: the type Celsius and Fahren is not the same type and cannot be compared.
 *
 * Data types
 *     basic, hybird, reference and interface.
 *     hybird    - array, struct
 *     reference - pointer, slice, dictionary, function, pipeline
 *
 * Const generator
 *     Sunday to Saturday will be initialized from 0 - 6.
 *     type weekday int
 *     const (
 *         Sunday weekday = itoa
 *         Monday
 *         Tuesday
 *         Wednesday
 *         Thurday
 *         Friday
 *         Saturday
 *     )
 *
 * Array, slice, map and struct
 *     [...] will count the size on initialization.
 *     var q [...]int = {1, 2, 3}
 *     var q [...]int = {0: 1, 1: 2, 2: 3}
 *
 * Map
 *     key -> value, key must be compared, value not.
 *     args := make(map[string]int)
 *     delete(map, key)
 *     the element of map cannot be referenced, like &args["Tom"]
 *     The iteration on map is unstable.
 *
 * Struct
 *     type Point struct {
 *         x, y int
 *     }
 *     Both instance and pointer use . to access members.
 *
 * Function
 *     func sum(vals ... int) int {
 *     }
 *     Which means it will accept any more than 1 args, the vals will treat as
 *     slice. If us slice as args. The function call should be
 *     values := int[]{1, 2, 3, 4}
 *     sum(values...)
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
	net_http_url_go_runtine()
	// net_web_server()

	flag_parser()
	print_valist_reuse()
	structBehavior()
	structEmbedded()
	fmt.Printf("pli28 panic %d\n", noReturn())
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
                /* return opened file and err */
		if fop, err := os.Open(arg); err != nil {
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

func fetch_url(url string, ch chan<- string ) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) /* send to channel ch */
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s, %d\n", url, err)
		return
	}

	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.4fs	%7d	%s", secs, nbytes, url);
}

func net_http_url_go_runtine() {
	url_list := []string{"https://golang.org", "http://gopl.io",
			"http://www.google.com", "http://www.bing.com"}
	ch := make(chan string)

	for _, url := range url_list {
		go fetch_url(url, ch) /* start a go rountine */
	}

	for i := 0; i < len(url_list); i++ {
		if is_print {
			/*
			 * Here main thread reach here, it will wait until
			 * first goruntine finish touch channel string. And
			 * then continue wait the second fast one to finish.
			 */
			fmt.Println(<-ch)
		}
	}
}

var mu sync.Mutex
var count int = 0

func web_server_handler (w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	fmt.Fprintf(w, "URL.Path = %q, count %d\n", r.URL.Path, count);
	mu.Unlock()
}

func net_web_server() {
	http.HandleFunc("/", web_server_handler)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func ftoc(f float64) float64 {
	return (f - 32) * 5 / 9;
}

func flag_parser() {
	/* name default-value descriptor, and flag return pointers */
	var n = flag.Bool("n", false, "omit trailing newline")
	var sep = flag.String("s", " ", "separator")
	flag.Parse()

	/* flag.Args return slice of string for all args */
	if is_print {
		fmt.Print(strings.Join(flag.Args(), *sep))
	}

	if !*n && is_print {
		fmt.Println()
	}
}

func print_valist_reuse() {
	i := 0666
	if is_print {
		/* [1] tell the Printf to reuse the first arg i */
		fmt.Printf("%d %[1]x %[1]X\n", i)
	}
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash + 1:]

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s;
}

func intToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')

	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}

	buf.WriteByte(']')
	return buf.String()
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1

	if zlen <= cap(x) {
		z = x[:zlen] /* here z and x will share the low level data */
	} else {
		zcap := zlen

		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}

		z = make([]int, zlen, zcap)
		copy(z, x)
	}

	z[len(x)] = y
	return z
}

func NonEmpty(str []string) []string {
	i := 0

	for _, s := range str {
		/*
                 * reuse the memory of given string slice, will overwrite low
                 * level memory.
                 */
		if s != "" {
			str[i] = s;
			i++;
		}
	}

	return str[:i];
}

func MapEqual(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}

	return true
}

type Point struct {
	x, y int
	Name string
}

type Wheel struct {
	Point
	length int
}

func structBehavior() {
	p := Point{0, 1, "null"}
	fmt.Println(p)
	/*
         * Another package cannot touch current package struct Point's internal
         * member x, y. They can only see Name here.
         */
}

func structEmbedded() {
	w := Wheel{Point{x:3, y:4, Name:"pli28"}, 20}

	w.x += w.y
	w.y -= 2
	fmt.Println(w)
}

func handleEOF() {
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Errorf("Read failed %v\n", err)
			break
		}
		fmt.Println(r)
	}
}

func squares() func() int {
	var x int
	/*
	 * Anonymous function will referece to variable x, foreach function call
	 * to anonymous function, it will keep the reference to seuares function
	 * variable x. For all, if there is a reference to anonymous function,
	 * the memory stack/hep of function squares will be keepped, so the
	 * anonymous function can access the variable and resource. Always, we 
	 * call it closure.
	 */
	return func() int {
		x++
		return x * x
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}

func deferClose(filename string) ([]byte, error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer f.Close()
	return ioutil.ReadAll(f)
}

func noReturn() (out int) {
	defer func() {
		if p := recover(); p != nil {
			out = p.(int)
		}
	}()

	panic(3)
}

