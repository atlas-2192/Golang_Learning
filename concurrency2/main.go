package main

import (
	"fmt"
	// "time"
)

type CheckMapOut struct {
	Result  string `json:"result"`
	Message string `json:"message"`
	MapUrl  string `json:"url"`
}

func output(signal chan int) {
	fmt.Println("In output")

	str := <-signal
	fmt.Printf("get channel %s", str)
}

func main() {
	fmt.Println("start")

	c1 := make(chan int, 1)
	go output(c1)

	c1 <- "sss"
}
