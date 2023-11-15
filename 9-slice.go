package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0])

	// with slice operation
	var newFruits = fruits[0:2]
	fmt.Println(newFruits)
}
