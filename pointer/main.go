package main

import "fmt"

func main() {
	var a []int
	var b []int
	a = append(a, 0)
	b = append(b, 0)
	p := &a[0]
	fmt.Printf("a[0] = %d pointer=%d, p = %d \n", a[0], &a[0], *p)
	a[0] = 2
	fmt.Printf("a[0] = %d pointer=%d, p = %d \n", a[0], &a[0], *p)
	/*
		a[0] = 0, p = 0
		a[0] = 2, p = 2
	*/
	var c []int
	var d []int
	c = append(c, 0)
	d = append(d, 0)
	p2 := &c[0]
	fmt.Printf("c[0]=%d pointer=%d, p2 = %d\n", c[0], &c[0], *p2)
	c = append(c, 1)
	c[0] = 2
	fmt.Printf("c[0]=%d pointer=%d, p2 = %d\n", c[0], &c[0], *p2)
	/*
		c[0]=0, p2 = 0
		c[0]=2, p2 = 0
	*/
}
