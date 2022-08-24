// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/rogpeppe/go-internal/par"
)

func main() {
	var w par.Work
	const N = 100
	for i := 0; i < N; i++ {
		w.Add(i)
	}
	start := time.Now()
	var n int32
	w.Do(N, func(x interface{}) {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(n)
		atomic.AddInt32(&n, +1)
	})
	if n != N {
		fmt.Println("par.Work.Do did not do all the work")
	}
	if time.Since(start) < N/2*time.Millisecond {
		fmt.Println("done!")
		return
	}
}
