package main

import (
	"fmt"
	"time"
)

func counter(c chan int) {
	for i := 0; ; i++ {
		c <- i
	}
}

func terminator(c chan string) {
	// terminate counter after 2 sec
	time.Sleep(time.Second * 2)
	c <- "EOD" // termination signal
}

func main() {
	c := make(chan int)
	c1 := make(chan string)

	go counter(c)
	go terminator(c1)

loop:
	for {
		time.Sleep(time.Second * 1)
		select {
		case msg := <-c:
			fmt.Println("counter value ", msg)
		case msg := <-c1:
			fmt.Println("Terminating counter ", msg)
			break loop
		default:
			fmt.Println("Nothing........")
		}
	}

	var input string
	fmt.Scanln(&input)
}
