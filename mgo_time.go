package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Person struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string
	Phone     string
	Timestamp time.Time
}

var (
	IsDrop = false
)

func main() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	// Drop Database
	if IsDrop {
		err = session.DB("test").DropDatabase()
		if err != nil {
			panic(err)
		}
	}

	// Collection People
	c := session.DB("MongoTest1").C("test")

	// Index
	index := mgo.Index{
		Key:        []string{"name", "phone"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	// Insert Datas
	err = c.Insert(&Person{Name: "fff", Phone: "+55 53 1234 4321", Timestamp: time.Now().AddDate(0, 0, -3)},
		&Person{Name: "gggg", Phone: "+66 33 1234 5678", Timestamp: time.Now().AddDate(0, 0, 4)})

	if err != nil {
		panic(err)
	}

	// Query One
	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).Select(bson.M{"phone": 0}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Phone", result)

	// Query All
	var results []Person
	err = c.Find(bson.M{"name": bson.M{"$exists": true}}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ")
	for _, v := range results {
		fmt.Println(v.Name, ":", v.Timestamp.Format("2006-01-02T15:04:05Z07:00"))
	}

	// Update
	colQuerier := bson.M{"name": "Ale"}
	change := bson.M{"$set": bson.M{"phone": "+86 99 8888 7777", "timestamp": time.Now()}}
	err = c.Update(colQuerier, change)
	if err != nil {
		panic(err)
	}

	// Query All
	err = c.Find(bson.M{"name": bson.M{"$exists": true}}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ")
	for _, v := range results {
		fmt.Println(v)
	}
}
