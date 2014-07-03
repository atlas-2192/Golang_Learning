package main

import (
	"fmt"
	"time"
)

func initX() {
	fmt.Print("init")

}
func main() {
	fmt.Print("hello world")
	initX()
	var t = time.Now().Unix()
	//var t2 = time.Now()
	fmt.Print(t)
	//fmt.Print(t2)
}
