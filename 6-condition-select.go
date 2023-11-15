package main

import "fmt"

func main() {
	// if, else if & else keyword
	var myPoint = 8

	if myPoint == 10 {
		fmt.Println("istimewa")
	} else if myPoint >= 6 {
		fmt.Println("lulus")
	} else {
		fmt.Printf("(%d), tidak lulus", myPoint)
	}

	// variable temporary if - else
	var ourPoint = 8840.0

	if percent := ourPoint / 100; percent >= 100 {
		fmt.Printf("%.1f %s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f %s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f %s not bad\n", percent, "%")
	}

	// switch keyword
	var justPoint = 6

	switch justPoint {
	case 10, 9, 8:
		fmt.Println("perfecto!")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	var otherPoint = 6
	switch {
	case otherPoint == 8:
		fmt.Println("perfect")
	case (otherPoint < 8) && (otherPoint > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("bruhh//")
			fmt.Println("you need to learn more")
		}
	}

	// fallthrough switch
	var pointx = 6

	switch {
	case pointx == 8:
		fmt.Println("perfect")
	case (pointx < 8) && (pointx > 3):
		fmt.Println("awesome")
		fallthrough
	case pointx < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}
}
