package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type book struct {
	_id         string
	isbn        string
	title       string
	releaseDate string
	listPrice   int
	pubId       string
}

func main() {
	fmt.Println("Starting connect mongoDB....")
	session, err := mgo.Dial(SERVER_CONNECT_STRING)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	fmt.Println("Connection works....")
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("MongoTest1").C("books")

	fmt.Println("Getting  Database count....")
	count, err2 := c.Find(bson.M{} /*bson.M{"_id": 2}*/).Count() // .One(&result)
	if err2 != nil {
		panic(err)
	}
	fmt.Printf("count = %d\n", count)

	result := book{}
	fmt.Println("Getting  data....")
	err2 = c.Find(bson.M{"_id": "5"}).One(&result)
	if err2 != nil {
		panic(err)
	}

	fmt.Printf("result =%s\n", result)
}
