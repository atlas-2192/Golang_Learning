package main

import "code.google.com/p/go-tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk_acc(t, ch)
	close(ch)
}

func walk_acc(t *tree.Tree, ch chan int) {
	if t != nil {
		walk_acc(t.Left, ch)
		ch <- t.Value
		//fmt.Println(t.Value)
		walk_acc(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for v1 := range ch1 {
		fmt.Println(v1)
		if v1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	//tt := tree.New(1)
	//fmt.Println("out tree=", tt.String())
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}

	if Same(tree.New(1), tree.New(2)) {
		fmt.Println("is the sam")
	} else {
		fmt.Println("is not the same ")
	}

}
