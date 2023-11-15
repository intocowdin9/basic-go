package main

import (
	"errors"
	"fmt"
	"strings"
)

//
/*
func main() {
	var input string
	fmt.Print("type some number: ")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}
}
*/

// custom error
func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

func main() {
	var name string
	fmt.Print("type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		// fmt.Println(err.Error())
		// penggunaan panic
		panic(err.Error())
	}
}
