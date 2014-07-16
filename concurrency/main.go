package main

import "fmt"

func sum(a []int, c chan int) {
	fmt.Println("Arra is =", a)
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c1 := make(chan int)
	//a[:len(a)] ==
	//
	//Array[a:b]
	//remove a or get b
	//
	go sum(a[2:5], c1)
	fmt.Println("run first go sum")
	go sum(a[len(a)/2:], c1)
	fmt.Println("run second go sum")
	x, y := <-c1, <-c1 // receive from c
	// <-channel will let the goroutine start to work.

	fmt.Println(x, y, x+y)

	/*
		run first go sum
		run second go sum
		Arra is = [8 -9 4]
		Arra is = [-9 4 0]
		3 -5 -2
	*/

	//------------------------
	//Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channe
	//
	c2 := make(chan int, 3)
	c2 <- 2
	c2 <- 3
	c2 <- 1
	fmt.Println(cap(c2))
	fmt.Printf("1=%d, 2=%d = %d \n ", <-c2, <-c2, <-c2)
	//------------------------

	c3 := make(chan int, 10)
	go func(c chan int) {
		fmt.Println("in inner func")
		for i := 1; i < cap(c); i++ {
			fmt.Println(i)
			c <- i * i
		}
		close(c) //actually it only run 9 not 10, need close or it will failed on range in #61
	}(c3)

	for j := range c3 {
		fmt.Println("out func")
		fmt.Println(<-c3 * j)
	}
}
