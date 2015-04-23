package main

import (
	"encoding/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

type GoogleIAP struct {
	Kind               string `json:"kind"`
	PurchaseTimeMillis string `json:"purchaseTimeMillis"`
	PurchaseState      string `json:"purchaseState"`
	ConsumptionState   bool   `json:"consumptionState"`
	DeveloperPayload   string `json:"developerPayload"`
}

func main() {
	data, err := ioutil.ReadFile("downloaded.json")
	if err != nil {
		log.Fatal(err)
	}
	conf, err := google.JWTConfigFromJSON(data, "https://www.google.com/m8/androidpublisher")
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(oauth2.NoContext)
	res, err := client.Get("https://www.googleapis.com/androidpublisher/v2/applications/`com.example.iap`/purchases/products/`test`/tokens/`TOKEN`")
	if err != nil {
		log.Println("err:", err)
	}
	body, err := ioutil.ReadAll(res.Body)
	log.Println(string(body))

	appResult := &GoogleIAP{}

	err = json.Unmarshal(body, &appResult)
	log.Printf("Receipt return %+v \n", appResult)

	//To transfer to purchase time millisecond to time.Time
	time_duration, _ := strconv.ParseInt(appResult.PurchaseTimeMillis, 10, 64)
	log.Println(time_duration)

	time_purchase := time.Unix(time_duration/1000, 0)
	log.Println(time_purchase.Local())

	// Compare with receipt to make sure the receipt if valid.

	// 1. appResult.PurchaseTimeMillis need equal to purchaseTime get from App.
	// 2. appResult.ConsumptionState need to be 1. (User already consumed)
	// 3. appResult.PurchaseState need to be 0. (Order is completed.)

	//Do_SOMETHING_IN_SERVER_PART()

}
