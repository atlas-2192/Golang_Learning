package main

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
)

func main() {
	// Your credentials should be obtained from the Google
	// Developer Console (https://console.developers.google.com).
	// Navigate to your project, then see the "Credentials" page
	// under "APIs & Auth".
	// To create a service account client, click "Create new Client ID",
	// select "Service Account", and click "Create Client ID". A JSON
	// key file will then be downloaded to your computer.
	data, err := ioutil.ReadFile("downloaded.json")
	if err != nil {
		log.Fatal(err)
	}
	conf, err := google.JWTConfigFromJSON(data, "https://www.google.com/m8/feeds")
	if err != nil {
		log.Fatal(err)
	}
	// Initiate an http.Client. The following GET request will be
	// authorized and authenticated on the behalf of
	// your service account.
	client := conf.Client(oauth2.NoContext)
	res, err := client.Get("https://www.google.com/m8/feeds/contacts/default/full")
	if err != nil {
		log.Println("err:", err)
	}
	body, err := ioutil.ReadAll(res.Body)
	log.Println(string(body))
}
