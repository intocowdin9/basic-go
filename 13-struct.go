package main

import (
	"fmt"
)

type student struct {
	name  string
	grade int
}

func main() {
	// var s1 student
	// s1.name = "john wick"
	// s1.grade = 2

	// fmt.Println("name	:", s1.name)
	// fmt.Println("grade	:", s1.grade)

	// with inisialize object struct
	var s1 = student{}
	s1.name = "wick"
	s1.grade = 2

	var s2 = student{"ethan", 2}
	var s3 = student{name: "jason"}

	fmt.Println("student 1:", s1.name)
	fmt.Println("student 2:", s2.name)
	fmt.Println("student 3:", s3.name)

	// call embedded struct
	mainEmbeddedStruct()
	// call anonymous struct
	mainAnonymousStruct()
}

// embedded struct with same name property
type person struct {
	name string
	age  int
}

type teacher struct {
	person
	age   int
	grade int
}

func mainEmbeddedStruct() {
	var t1 = teacher{}
	t1.name = "wick"
	t1.age = 31
	t1.person.age = 32

	fmt.Println(t1.name)
	fmt.Println(t1.age)
	fmt.Println(t1.person.age)
}

// anonymous struct
func mainAnonymousStruct() {
	// anonymous struct without initialize
	var s1 = struct {
		person
		grade int
	}{}

	s1.person = person{"wick", 21}
	s1.grade = 2

	fmt.Println("name 	:", s1.person.name)
	fmt.Println("age 	:", s1.person.age)
	fmt.Println("grade	:", s1.grade)

	// anonymous struct with initialize
	var s2 = struct {
		person
		grade int
	}{
		person: person{"wick", 21},
		grade:  2,
	}

	s2.person = person{"john", 22}
	s2.grade = 3

	fmt.Println("name 	:", s2.person.name)
	fmt.Println("age 	:", s2.person.age)
	fmt.Println("grade 	:", s2.grade)
}
