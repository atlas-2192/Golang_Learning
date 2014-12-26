package main

import (
	"fmt"
	"github.com/koding/kite"
	"net/url"
)

func main() {
	k := kite.New("first", "1.0.0")
	k.Config.Port = 6000
	k.HandleFunc("square", func(r *kite.Request) (interface{}, error) {
		a := r.Args.One().MustFloat64()
		fmt.Println("get input %f", a)
		return a * a, nil
	})

	k.Register(&url.URL{Scheme: "http", Host: "localhost:6000/kite"})
	k.Run()
}
