package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	/*
		// fungsi time.Sleep()
		fmt.Println("start")
		time.Sleep(time.Second * 4)
		fmt.Println("after 4 second")

		// fungsi time.NewTimer()
		var timer = time.NewTimer(4 * time.Second)
		fmt.Println("start")
		<-timer.C
		fmt.Println("finish")

		// fungsi time.AfterFunc()
		var ch = make(chan bool)

		time.AfterFunc(4*time.Second, func() {
			fmt.Println("expired")
			ch <- true
		})

		fmt.Println("start")
		<-ch
		fmt.Println("finish")

		// fungsi time.After()
		<-time.After(4 * time.Second)
		fmt.Println("Expired")
	*/

	// call kombinasi timer & goroutine
	mainCombWithGoroutine()

}

func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntime out! no answer more than", timeout, "seconds")
	os.Exit(0)
}

func mainCombWithGoroutine() {
	var timeout = 5
	var ch = make(chan bool)

	go timer(timeout, ch)
	go watcher(timeout, ch)

	var input string
	fmt.Print("what is 725/25 ? ")
	fmt.Scan(&input)

	if input == "29" {
		fmt.Println("the answer is right!")
	} else {
		fmt.Println("the answer is wrong!")
	}
}
