package main

import "fmt"

func main() {
	//a := make([]int, 1, 1)
	var a []int
	//b := make([]int, 1, 1)
	var b []int

	// append works on nil slices.
	a = append(a, 0)
	b = append(b, 0)
	p := &a[0]

	// printSlice("a", a)
	// fmt.Printf("p = %d , point %d \n", *p, p)
	// fmt.Printf("a point = %d , point %d \n", &a, &a[0])
	// fmt.Printf("b point = %d , point %d \n", &b, &b[0])

	// // the slice grows as needed.
	a = append(a, 1)
	// printSlice("a", a)
	fmt.Printf("p = %d, point %d \n", *p, p)
	// fmt.Printf("point = %d , point %d , point %d \n", &a, &a[0], &a[1])
	a[0] = 2
	fmt.Printf("a[0] = 2 p = %d , point %d \n", *p, p)
	fmt.Printf("a[0] = 2 point = %d , point %d , point %d \n", &a, &a[0], &a[1])

	printSlice("a", a)
	printSlice("b", b)
	//loop 50 times to check result
	for i := 1; i <= 50; i++ {
		a = append(a, i+2)
	}
	fmt.Println("Add 50 times to over memory slot")
	fmt.Printf("a[0] = %d , point %d \n", a[0], &a[0])
	fmt.Printf("b[0] = %d , point %d \n", b[0], &b[0])
	fmt.Printf("p = %d , point %d \n", *p, p)
	a[0] = 3

	fmt.Printf("a[0] = %d , point %d \n", a[0], &a[0])
	fmt.Printf("b[0] = %d , point %d \n", b[0], &b[0])
	fmt.Printf("p = %d , point %d \n", *p, p)

	printSlice("a", a)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
