package main

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}

func connect() *mgo.Session {
	var session, err = mgo.Dial("localhost")
	if err != nil {
		os.Exit(0)
	}
	return session
}

func insert() {
	var session = connect()
	defer session.Close()
	var collection = session.DB("belajar_golang").C("student")

	var err = collection.Insert(&student{"Wick", 2}, &student{"Ethan", 2})
	if err != nil {
		fmt.Println(err.Error())
	}
}

func find() {
	var session = connect()
	defer session.Close()
	var collection = session.DB("belajar_golang").C("student")

	var result = student{}
	var selector = bson.M{"name": "Wick"}
	var err = collection.Find(selector).One(&result)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Name :", result.Name)
	fmt.Println("Grade :", result.Grade)
}

func update() {
	var session = connect()
	defer session.Close()
	var collection = session.DB("belajar_golang").C("student")

	var selector = bson.M{"name": "Wick"}
	var changes = student{"John Wick", 2}
	var err = collection.Update(selector, changes)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func remove() {
	var session = connect()
	defer session.Close()
	var collection = session.DB("belajar_golang").C("student")

	var selector = bson.M{"name": "John Wick"}
	var err = collection.Remove(selector)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	// insert()
	// find()
	update()
	remove()
}
