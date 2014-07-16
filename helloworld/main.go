package main

import (
	//"code.google.com/p/go-tour/pic"
	"fmt"
	"image"
	"image/color"
	"math"
	"net/http"
	"os"
	"pkg/stacker"
	"strings"
	"time"
)

func initX() {
	fmt.Print("init")

}

func foo2(in_x int) (x, y int) {
	x = in_x + 10
	y = in_x - 10
	return
}

const (
	big   = 1 << 100
	small = big >> 99
)

func WordCount(s string) map[string]int {
	split := strings.Fields(s)
	out := make(map[string]int)
	for i := 0; i < len(split); i++ {
		if out[split[i]] != 0 {
			out[split[i]] += 1
		} else {
			out[split[i]] = 1
		}
	}
	return out
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum = sum + x
		return sum
	}
}

func fibonacci() func() (ret int) {
	fib1 := 0
	fib2 := 1
	return func() (ret int) {
		ret = fib1 + fib2
		fib1 = fib2
		fib2 = ret
		return
	}
}

//similar with C++ class. Methonds
type Circle struct {
	radius float64
}

func (v *Circle) area() float64 {
	return math.Pow(v.radius, 2) * math.Pi
}

// other method
type squre_z struct {
	x float64
	y float64
}

func (sz squre_z) area() float64 {
	return sz.x * sz.y
}

type model interface {
	area() float64
}

//Duck Typing
type abstract_t interface {
	//Unlink python, golang need define func in interface (point class) first.
	foo_func()
}

type sub_t1 struct{}

func (t1 sub_t1) foo_func() {
	fmt.Print("sub_t1 foo \n")
}

type sub_t2 struct{}

func (t2 sub_t2) foo_func() {
	fmt.Print("sub_t2 foo \n")
}

func show_foo(t abstract_t) {
	t.foo_func()
}

//Goroutine
func foo_rountine(msg1 string, msg2 string) {
	go func(msg string) { //
		//time.Sleep(delay)
		fmt.Print("inner", msg)
	}(msg1) //must add parenthess
	fmt.Print(msg2)
}

//http
type HelloSrv struct{}

func (srv HelloSrv) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello server<br> test it")
}

//image
type Image struct {
	width, height int
}

func (ii Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, ii.width, ii.height)
}

func (ii Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (ii Image) At(x, y int) color.Color {
	c := uint8(x - y)
	return color.RGBA{c, c, 255, 255}
}

func summ(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	fmt.Print("hello world")
	fmt.Print("\n")
	initX()
	var t = time.Now().Unix()
	fmt.Print(t)
	fmt.Print("\n")

	//Arrary maniplation
	var ary [10]int
	ary[0] = 5
	fmt.Printf("array is %d \n", ary[0])
	a := [3]int{3, 5, 7}
	fmt.Printf("other ary is %d\n", a[2])
	who := "string Woo\n"
	fmt.Print(who)

	if len(os.Args) > 1 {
		who = strings.Join(os.Args[1:], "")
	}
	fmt.Print("hello", who)

	var haystack stack.Stack
	haystack.Push("text")
	haystack.Push(14)
	item, error := haystack.Pop()
	if error == nil {
		fmt.Print("Pop out ->", item)
	}

	//About Map
	m1 := make(map[string]string)
	m1["index1"] = "value1"
	m2 := m1
	m2["index1"] = "value2"
	fmt.Print("\nmap out ->", m2["index1"], "value2 =", m2["index2"])

	//function return as it named
	x, y := foo2(20)
	fmt.Print("x=%d,y=%d \n", x, y)
	//using constant
	fmt.Print("big=", small, "\n")

	//for loop
	for i := 0; i <= 5; i++ {
		fmt.Print(".")
	}

	sum := 1
	for sum < 10 {
		sum += 1
		fmt.Print("o")
	}

	//for range for slice. use nil feature. for could easy transfer to while via interaction
	var pow = []int{1, 2, 4, 8, 16}
	for i, v := range pow {
		fmt.Printf("i=%d, v=%d \n", i, v)
	}

	fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))

	//WordCount tester
	fmt.Print(WordCount("I ate a donut. Then I ate another donut."), "\n")

	//function closure about using function as a value
	fx, fy := adder(), adder()
	fmt.Printf("output fx(3)=%d\n", fx(3))
	fmt.Printf("output fx(3)=%d\n", fx(3)) //value will keep 3+3
	fmt.Printf("output fy(10)=%d\n", fy(10))

	//fibonnaci
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	//switch case
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Print("today!\n")
	case today + 1:
		fmt.Print("tomorrow\n")
	}

	//using Methods
	c1 := Circle{4}
	fmt.Printf("area is %f\n", c1.area())

	//Interface like parent class point.
	var m model
	m = &c1
	fmt.Print(m.area())

	//duck typing
	var tt abstract_t //declare interface
	t1 := sub_t1{}    //new instance struct sub_t1
	tt = t1           //assign value and its func to interface. (c++ func point assignment)
	show_foo(tt)

	//goroutine
	foo_rountine("in routine\n", "out routine\n")

	//very slow, actually if no input it will goto end without run goroutine
	go func(msg string) { //
		//time.Sleep(delay)
		fmt.Print("inner", msg)
	}("msg1")

	//very slow, actually if no input it will goto end without run goroutine
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	show_foo(tt)
	show_foo(tt)
	show_foo(tt)

	//Wait for goroutine, only works in execution file
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")

	//http
	//fmt.Println("init http server")
	//var v HelloSrv
	//http.ListenAndServe("localhost:4000", v)
	//fmt.Println("finished")

	//image
	//mm := Image{255, 255}
	//pic.ShowImage(mm)

	//channel part
	aa := []int{7, 2, 8, -9, 4, 0}
	ch := make(chan int)
	summ(aa[:len(aa)/2], ch)
}
