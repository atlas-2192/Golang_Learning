package main

import "fmt"

func main() {
	var a []int
	var b []int
	a = append(a, 0)
	b = append(b, 0)
	p := &a[0]
	fmt.Printf("a[0] = %d, p = %d \n", a[0], *p)

	// the slice grows as needed.
	a[0] = 2
	fmt.Printf("a[0] = %d, p = %d \n", a[0], *p)

	var c []int
	var d []int

	c = append(c, 0)
	d = append(d, 0)

	p2 := &c[0]
	c = append(c, 1)
	// printSlice("a", a)
	// fmt.Printf("p = %d, point %d \n", *p, p)
	// fmt.Printf("point = %d , point %d , point %d \n", &a, &a[0], &a[1])
	fmt.Printf("c[0]=%d, p2 = %d\n", c[0], *p2)
	c[0] = 2
	fmt.Printf("c[0]=%d, p2 = %d\n", c[0], *p2)
	// fmt.Printf("c[0] = 2 p2 = %d , point %d \n", *p2, p2)
	// fmt.Printf("c[0] = 2 point = %d , point %d , point %d \n", &c, &c[0], &c[1])

}
