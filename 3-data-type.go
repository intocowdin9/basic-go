package main

import "fmt"

func main() {
	// non-decimal
	var positiveNumber uint8 = 89
	var negativeNumber = -1234567799

	fmt.Printf("positive number\t: %d\n", positiveNumber)
	fmt.Printf("negative number\t: %d\n", negativeNumber)

	// decimal
	var decimalNumber = 2.70
	fmt.Printf("decimal number\t: %f\n", decimalNumber)
	fmt.Printf("decimal number\t: %.3f\n", decimalNumber)

	// boolean
	var exist bool = true
	fmt.Printf("exist? %t\n", exist)

	// string
	var message string = `"welcome".
	"and good luck!"`
	fmt.Println("\n", message)
}