package library

import "fmt"

func SayHello() {
	fmt.Println("hello")
	introduce("Rafli")
}

func introduce(name string) {
	fmt.Println("my name is: ", name)
}

// function init
var Student = struct {
	Name  string
	Grade int
}{}

func init() {
	Student.Name = "John Wick"
	Student.Grade = 2

	fmt.Println("--> library/library.go imported")
}
