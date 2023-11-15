package main

import (
	"fmt"
	"strings"
)

type student struct {
	name  string
	grade int
}

func (s student) sayHello() {
	fmt.Println("hello", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-2]
}

func main() {
	var s1 = student{"john wick", 21}
	s1.sayHello()

	var name = s1.getNameAt(2)
	fmt.Println("nama panggilan : ", name)
}

/**
// method pointer
func (s *student) sayHello() {
	fmt.Println("halo", s.name)
}
func main() {
	var s1 = student{"john wick", 21}
	s1.sayHello()
}

func mainMethodPointer() {
	// pengaksesan method dari variabel objek biasa
	var s1 = student{"john wick", 21}
	s1.sayHello()
	// pengaksesan method dari variabel objek pointer
	var s2 = &student{"ethan hunt", 22}
	s2.sayHello()
}
**/
