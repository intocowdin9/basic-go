package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["january"] = 1
	chicken["february"] = 2

	fmt.Println("January", chicken["january"])
	fmt.Println("Mei", chicken["mei"])

	// with initialize values map
	var chicken1 = map[string]int{"jan": 10, "feb": 20}
	fmt.Println(chicken1["jan"])

	// delete item map
	delete(chicken1, "jan")

	// with for - range
	for key, val := range chicken1 {
		// delete
		fmt.Println(key, " \t", val)
	}

	// detection of item
	var value, isExist = chicken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exits")
	}
}
