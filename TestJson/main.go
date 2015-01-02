package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type test_struct struct {
	Test string
}

func test(rw http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read error")
	}

	var dat map[string]interface{}

	//Data: `{"name":"test1", "num":"2", "data": [{"data1": "value1", "data2": "value2" }] }`
	err = json.Unmarshal(body, &dat)
	if err != nil {
		fmt.Println("Unmarshal error")
	}
	fmt.Println(dat)

	for key, value := range dat {
		fmt.Println("Map1 Key:", key, "Value:", value)
	}

	map_body := dat["data"].([]interface{})
	fmt.Println(map_body)
	var dat2 map[string]interface{}
	// Get first element in array which is a map.
	dat2 = map_body[0].(map[string]interface{})
	fmt.Println(dat2)

	for key, value := range dat2 {
		fmt.Println("Map2 Key:", key, "Value:", value)
	}
}

func main() {
	http.HandleFunc("/test", test)
	fmt.Println("Server Start..")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
