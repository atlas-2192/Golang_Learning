package main

import "fmt"

func gen(nums ...int) chan int {
	out := make(chan int)
	fmt.Println("enter gen")
	go func() {
		fmt.Println("enter go gen")
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	fmt.Println("leave gen")
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	fmt.Println("enter sq")
	go func() {
		fmt.Println("enter go sq")
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	fmt.Println("leave sq")
	return out
}

func main() {
	// Set up the pipeline.
	c := gen(2, 3)
	out := sq(c)

	// Consume the output.
	fmt.Println("start request number")
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9

	//another way
	for n := range sq(sq(gen(4, 5))) {
		fmt.Println("Another result = ", n)
	}
}
