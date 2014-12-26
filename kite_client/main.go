package main

import (
	"fmt"
	"github.com/koding/kite"
	"github.com/koding/kite/protocol"
)

func main() {
	k := kite.New("second", "1.0.0")

	kites, _ := k.GetKites(&protocol.KontrolQuery{
		Username:    k.Config.Username,
		Environment: k.Config.Environment,
		Name:        "first",
	})

	if kites != nil {
		client := kites[0]
		client.Dial()

		//default API which return "pong"
		//response, _ := client.Tell("kite.ping")
		//fmt.Println(response.MustString())

		response2, _ := client.Tell("square", 4)
		fmt.Println(response2.MustFloat64())
	} else {
		fmt.Println("Can not get any RPC server")
	}
}
