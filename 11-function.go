package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var names = []string{"John", "Wick"}
	printMessage("hello, ", names)

	// call with return value
	withReturnValue()
	// call with process function
	withStopProcess()
	// call function multiple return values
	var diameter float64 = 15
	var area, circumference = calculate(diameter)
	fmt.Printf("luas lingkaran\t\t: %.2f\n", area)
	fmt.Printf("keliling lingkaran\t\t: %.2f\n", circumference)
	// call function variodic
	mainFunctionVariodic()
	// call function common parameter & variadic
	yourHobbies("wick", "sleeping", "eating")
	// call closure saving as var
	closureSavingAsVar()
	// call iife
	IIFESample()
	// call closure as return value
	mainFindMax()
	// contoh penggunaan fungsi string.Contains()
	var result = strings.Contains("Golang", "ang")
	fmt.Println("\n", result)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

// with return value
func withReturnValue() {
	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("\nrandom number :", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number :", randomValue)
}

func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}

// return for stoping process in function
func withStopProcess() {
	divideNumber(20, 2)
	divideNumber(20, 0)
	divideNumber(20, -9)
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d \n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}

// function multiple return
func calculate(d float64) (float64, float64) {
	var area = math.Pi * math.Pow(d/2, 2)
	var circumference = math.Pi * d

	return area, circumference
}

// function variadic
func mainFunctionVariodic() {
	var avg = calculateFunc(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("\nrata-rata : %.2f", avg)
	fmt.Println(msg)
	// other way
	// fmt.Printf("rata-rata : %.2f", avg)
}

func calculateFunc(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

// with common parameter & variadic
func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("\nhello my name is: %s\n", name)
	fmt.Printf("my hobbies are: %s\n", hobbiesAsString)
}

// closure saving as variable
func closureSavingAsVar() {
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	fmt.Printf("\ndata : %v\nmin : %v\nmax : %v\n", numbers, min, max)
}

// Immediately-Invoked Function Expression- IIFE (also closure)
func IIFESample() {
	var numbers = []int{2, 3, 4, 6, 4, 4, 4, 2, 2, 3, 3, 4, 2, 1, 10, 23}

	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(4)

	fmt.Println("\noriginal number :", numbers)
	fmt.Println("filtered number :", newNumbers)
}

// closure as return value
func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}

func mainFindMax() {
	var max = 3
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var howMany, getNumbers = findMax(numbers, max)
	var theNumbers = getNumbers()

	fmt.Println("\nnumbers\t:", numbers)
	fmt.Printf("find \t: %d\n", max)

	fmt.Println("found \t:", howMany)
	fmt.Println("value \t:", theNumbers)
}

// function as parameter
func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func mainFilter() {
	var data = []string{"wich", "jason", "ethan"}
	var dataContains0 = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})

	var dataLenght5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t\t:", data)
	// data asli : [wick jason ethan]
	fmt.Println("filter ada huruf \"i\"\t:", dataContains0)
	// filter ada huruf "i" : [wick]
	fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)
	// filter jumlah huruf "5"	: [jason ethan]
}
