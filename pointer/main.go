package main

import "fmt"

func main() {
	var a []int
	var b []int
	fmt.Printf("a len=%d, cap=%d \n", len(a), cap(a))
	a = append(a, 0)
	b = append(b, 0)
	p := &a[0]

	fmt.Printf("a[0] = %d pointer=%d, len=%d, cap=%d, p = %d \n", a[0], &a[0], len(a), cap(a), *p)
	a[0] = 2
	fmt.Printf("a[0] = %d pointer=%d, len=%d, cap=%d, p = %d \n", a[0], &a[0], len(a), cap(a), *p)
	/*
		a[0] = 0, p = 0
		a[0] = 2, p = 2
	*/
	var c []int
	var d []int
	fmt.Printf("c len=%d, cap=%d \n", len(c), cap(c))
	c = append(c, 0)
	d = append(d, 0)
	p2 := &c[0]
	fmt.Printf("a[0] = %d pointer=%d, len=%d, cap=%d, p2 = %d \n", c[0], &c[0], len(c), cap(c), *p2)
	c = append(c, 1)
	c[0] = 2
	fmt.Printf("a[0] = %d pointer=%d, len=%d, cap=%d, p2 = %d \n", c[0], &c[0], len(c), cap(c), *p2)

	/*
			c[0]=0, p2 = 0
			c[0]=2, p2 = 0

		  in http://play.golang.org/p/DuClFZt2cK, the cap increasement by machine dependency.
			c[0]=0, p2 = 0
			c[0]=2, p2 = *2*

	*/
}
