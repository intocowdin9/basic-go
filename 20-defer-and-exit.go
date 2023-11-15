package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("halo")
	// penerapan funsi os.Exit()
	os.Exit(1)
	fmt.Println("selamat datang")

}
