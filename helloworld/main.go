package main

import (
	"fmt"
	"os"
	"stacker"
	"strings"
	"time"
)

func initX() {
	fmt.Print("init")

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
}
