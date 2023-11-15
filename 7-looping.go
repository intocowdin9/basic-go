package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// with condition
	var i = 0

	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	// for without argument
	var ii = 0
	for {
		fmt.Println("Angka", ii)
		ii++
		if ii == 5 {
			break
		}
	}

	// with continue and break
	for iii := 1; iii <= 10; iii++ {
		if i%2 == 1 {
			continue
		}

		if iii > 8 {
			break
		}

		fmt.Println("Angka", iii)
	}

	// nested looping
	for iv := 0; iv <= 5; iv++ {
		for j := iv; j <= 5; j++ {
			fmt.Print(iv, " ")
		}
		fmt.Println()
	}

	// using label in looping
outerLoop:
	for v := 0; v < 5; v++ {
		for jj := 0; jj < 5; jj++ {
			if v == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", v, "][", jj, "]", "\n")
		}
	}
}
