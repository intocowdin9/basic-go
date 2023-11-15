package main

import "fmt"

func main() {
	// pointer
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numbreA (value)   :", numberA)
	fmt.Println("numbreA (address) :", &numberA)
	fmt.Println("numbreB (value)   :", *numberB)
	fmt.Println("numbreB (address) :", numberB)

	fmt.Println("")

	numberA = 5

	fmt.Println("numberA (value)  :", numberA)
	fmt.Println("numberA (address):", &numberA)
	fmt.Println("numberB (value) :", *numberB)
	fmt.Println("numberB (address):", numberB)

	// call pointer parameter
	mainPointerParam()

}

// Pointer Parameter
func mainPointerParam() {
	var number = 4
	fmt.Println("\nbefore :", number)

	change(&number, 10)
	fmt.Println("after :", number)
}

func change(original *int, value int) {
	*original = value
}
