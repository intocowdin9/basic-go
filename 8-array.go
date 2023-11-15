package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])

	var fruits = [4]string{"kurma", "apple", "melon"}

	fmt.Println("sum of elements \t\t", len(fruits))
	fmt.Println("all elements is \t", fruits)

	// Initilize initial value of array without sum elements
	var numbers = [...]int{1, 2, 3}

	fmt.Println("data array \t", numbers)
	fmt.Println("all element \t", len(numbers))

	// multi-demensional array
	// var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers1 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1 : ", numbers1)
	fmt.Println("numbers2 : ", numbers2)

	// loop with for
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("element %d : %s\n", i, fruits[i])
	}

	// loop with range and _
	for _, fruit := range fruits {
		fmt.Printf("name fruit : %s \n", fruit)
	}
}
