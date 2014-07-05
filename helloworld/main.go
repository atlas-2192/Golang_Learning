package main

import (
	"fmt"
	"math"
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

	var m model
	m = &c1
	fmt.Print(m.area())
}
