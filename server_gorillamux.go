package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type Payload struct {
	Staff Data
}

type Data struct {
	Fruit   Fruits
	Vaggies Vagetables
}

type Fruits map[string]int
type Vagetables map[string]int

func fabnacci(sum chan int, quit chan int) (ret int) {
	x, y := 0, 1
	for {
		select {
		case sum <- x:
			x, y = y, x+y
			/* is rquivalant with follow
			z := x + y
			x = y
			y = z
			*/
		case <-quit:
			return
		}
	}
}

func main() {
	//init gorilla/mux
	m := mux.NewRouter()

	// Get all lists.
	m.HandleFunc("/", GetAllLists).Methods("GET")

	// route http to mux
	http.Handle("/", m)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func serveRest(res http.ResponseWriter, req *http.Request) {
	//fmt.Fprintln(res, "hello, world")
	// fmt.Fprintln(res, "Init concurrency")
	// fab := make(chan int)
	// quit := make(chan int)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Fprintln(res, "fab =", <-fab)
	// 	}
	// 	quit <- 0
	// }()
	// fabnacci(fab, quit)

	//Try json response
	// response, err := getJSONResponse()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Fprint(res, string(response))
}

func GetAllLists(w http.ResponseWriter, r *http.Request) {
	response, err := getJSONResponse()
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, string(response))
}
func getJSONResponse() ([]byte, error) {
	fruits := make(map[string]int)
	fruits["Apple"] = 2
	fruits["Tomato"] = 3

	vagetables := make(map[string]int)
	vagetables["Pappers"] = 5

	d := Data{fruits, vagetables}
	p := Payload{d}
	return json.MarshalIndent(p, "", " ")
}
