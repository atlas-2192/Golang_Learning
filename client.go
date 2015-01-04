package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
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

func main() {
	url := "http://localhost:5000/test"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{"name":"test1", "num":"2", "data": [{"data1": "value1", "data2": "value2" }] }`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
