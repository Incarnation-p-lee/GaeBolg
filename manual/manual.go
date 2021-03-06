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
 * 66. Embedded struct method can also be accessed by outlayer struct even if it do not implemented.
 * 67. Method also has value, like function. Page 222.
 * 68. Packaging means method cannot be touched by invoker.
 * 69. Interface can also be embedded, interface is null type, or any type, which contains no method.
 * 70. Interface also has values, one type and its value. For example, var w io.Writer.
 *     var w io.Writer
 *     +-----+
 *     | nil | <- type
 *     +-----+
 *     | nil | <- value
 *     +-----+
 *     w = os.Stdout
 *     +---------------------+
 *     | *os.File            | <- type
 *     +---------------------+
 *     | Pointer to os.File  | <- value
 *     +---------------------+
 *     So interface value always called dynamic type and its value, value also can be type descriptor.
 *     And value can be any big.
 *     var x interface{} = time.Now()
 *     +---------------------+
 *     | time.Time           | <- type
 *     +---------------------+
 *     | sec:  xxxxxxxxxxxx  |
 *     | usec: xxxxxxxxxxxx  | <- value
 *     | loc:  "UTC"         |
 *     +---------------------+
 * 71. Type assert, which is for interface value, looks like x.(T). T is type name.
 *     If T is not interface, it will return x's value on success, or panic
 *     If T is interface, it will check type T will fit type x, and return the interface
 *     with same type and value. It will change the interface func set.
 *     If interface type assert with 2 return value, it will no panic if failed.
 *     var w io.Writer = os.Stdout
 *     f, ok := w.(*os.File)
 * 72. For interface hold variables, one switch case can avoiding if-else if-else if
 *     switch x.(type) {
 *         case nil:
 *         case int, uint:
 *         case bool:
 *         case string:
 *         default:
 *     }
 * 73. If more than 2 types implement interface makes interface valuable, in case there
 *     must be someting can be abstracted and give interface higher level meanings.
 * 74. Gorounine is one concurrent execution unit, when main goruntime returns, all sub goruntime will
 *     be interrupt and exit.
 *     go func_call()
 * 75. Channel is the commnunication between goruntines. A chan has one send/recevie type.
 *     ch := make(chan int, capacity), as ch is a pointer like value. Channel has 2 operations, send
 *     and receive.
 *     ch <- x     // send to chan
 *     x = <- ch   // recevie from chan
 *     <- ch       // recevie can be discard
 *     chan also can be closed, then any send/recevie will result in panic.
 *     channel without buf send data will block, until other goruninte receive.
 *     receive also can return 2 values, (v, ok), so we can check if channel is closed.
 *     channel can be ranged
 *     naturals := make(chan int)
 *     squares := make(chan int)
 *     go func() {
 *         for x := 0; x < 100; x++ {
 *             naturals <- x
 *         }
 *         close(naturals)
 *     }()
 *
 *     go func() [
 *         for x := range naturals [
 *             squares <- x * x
 *         }
 *         close(squares)
 *     }()
 *     chan<- int means a channel only for send.
 *     <-chan int means a channel only for receive.
 *     The bufferred channel buffer looks like a queue.
 *     After chan close, we can use range to touch the pending data.
 * 76. Select for multi channel trigger, mostly for sync with multi channels.
 *     Any of select case triggerred will end the blocking of select statement.
 *     If chan abort has value it will receive or do nothing, the following recevie
 *     will not block. Default means if for now other cases are not ready, then execute default.
 *     select {
 *         case <- abort:
 *             fmt.Println
 *         default:
 *             // do nothing
 *     }
 * 77. Sync.Once, can be consider as one mutex and bool for initializer complete or not.
 *     var loadIconOnce sync.Once
 *     var icons map[string]image.Image
 *     func Icon(name string) image.Image {
 *          loadIconOnice.Do(loadIcons)
 *          return icons[name]
 *     }
 * 78. Os thread may hold one stack for each thread, but goruntine stack begins with 2K
 *     at first, and automatically increase at most 1GB. Also go has it own scheduler,
 *     some function like sleep, mutex will be scheduled. Go use m:n scheduler, means
 *     n thread for m goruntine. These sleep and mutex will not trap to kernel.
 * 79. Goruntine has no id.
 * 80. Package name is global unique string, package default name used as the package identifier.
 *         import "math/rand"
 *     var r = rand.Int()
 *     source file name with suffix test.go will be built by go test.
 *     Also, package can be renamed for conflict, only works in file scope.
 *         import "crypto/rand"
 *         import mrand "math/rand"
 *     And anonymous package import is allowed.
 *         import _ "math/rand"
 * 81. GOPATH specific the current go work dir, GOROOT for go install dir. For convention,
 *     one package, one directory.
 * 
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
 * Method
 *     type Point struct {
 *         x, y float64
 *     }
 *     func (p Point) Distance(q Point) float 64 {
 *         return xxx
 *     }
 *     func (p *Point) Distance(q Point) float 64 {
 *         return xxx
 *     }
 *     The p Point called Receiver.
 *     The method of embedded struct also embedded, the type ColoredPoint can
 *     also be the receiver of type Point.
 *     type ColoredPoint struct {
 *         Point
 *         Color color.RGBA
 *     }
 *
 * Interface
 *     io.Writer defines the rule of caller and callee, which makes any type implemented
 *     io.Writer on Fprintf works well.
 *     type Writer interface {
 *         Write(p []byte)(n int, err error)
 *     }
 *     func Fprintf(w io.Writer, format string, args ... interface{}) (int, err)
 *     func Printf(format string, args ... interface{}) (int, err) {
 *         return Fprintf(os.Stdout, format, args ...)
 *     }
 *     func Sprintf(format string, args ... interface{}) string {
 *         var buf bytes.Buffer
 *         Fprintf(&buf, format, args...)
 *         return buf.String()
 *     }
 *     var w io.Writer
 *     w = os.Stdout
 *     w = new(bytes.Buffer)
 *     w = time.Second       // Compiler error, time.Duration lack Writer method.
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
	goruntine()
	channelOnClose()
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

type FPoint struct {
	X, Y float64
}

type Path []FPoint

func (p FPoint) Add(q FPoint) FPoint {
	return FPoint{p.X + q.X, p.Y + q.Y}
}

func (p FPoint) Sub(q FPoint) FPoint {
	return FPoint{p.X - q.X, p.Y - q.Y}
}

func (path Path) methodValue(offset FPoint, is_add bool) {
	var op func(p, q FPoint) FPoint

	if is_add {
		op = FPoint.Add
	} else {
		op = FPoint.Sub
	}

	for i := range path {
		path[i] = op(path[i], offset)
	}
}

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x / 64, uint(x % 64)
	return word < len(s.words) && (s.words[word] & (1 << bit)) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x / 64, uint(x % 64)

	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}

	s.words[word] |= 1 << bit
}

func writeString(w io.Writer, s string) (n int, err error) {
	type stringWriter interface {
		WriteString(string)(n int, err error)
	}

	if sw, ok := w.(stringWriter); ok {
		/*
		 * use type assert to check if w implmented no extra memory
		 * WriteString.
		 */
		return sw.WriteString(s)
	}

	/* []byte() will allocate another chunk of memory */
	return w.Write([]byte(s))
}

func goruntine() {
	go spinner(100 * time.Millisecond)
	const n = 20
	fib := fibnumber(n)
	fmt.Println(fib)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fibnumber(x int) int {
	if x < 2 {
		return x
	}

	return fibnumber(x - 1) + fibnumber(x - 2)
}

func channelOnClose() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	out := make(chan int)
	var wg sync.WaitGroup

	for _, v := range data {
		wg.Add(1)

		go func(d int) {
			out <- d
			defer wg.Done()
		}(v)
	}

	go func() {
		// fmt.Printf("Waiting goruntimes")
		wg.Wait()
		close(out)
	}()

	/* after chan close the data still here */
	var total int
	for v := range out {
		total += v
	}

	fmt.Println(total)
}

