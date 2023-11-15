package main

import "fmt"

func main() {
	// op-aritmatika
	var value = (((2+6)%3)*4 - 2) / 3
	// op-perbandingan
	var isEqual = (value == 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	// op-logika
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Println("left && right \t:", leftAndRight)

	var leftOrRight = left || right
	fmt.Println("left || right \t:", leftOrRight)

	var leftReverse = !left
	fmt.Println("!left \t:", leftReverse)
}
