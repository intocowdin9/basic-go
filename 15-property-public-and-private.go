package main

import (
	"belajar-golang-level-akses/library"
	"fmt"
)

func main() {
	// publik akses diawali huruf besar
	library.SayHello()
	// private akses diawali huruf kecil
	// library.introduce("rafli") // error karna tidak bisa diakses

	// with call library init(inside)
	fmt.Printf("name  : %s\n", library.Student.Name)
	fmt.Printf("grade : %d\n", library.Student.Grade)
}
