package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go func() {
		// time.After(4 * time.Second)
		time.Sleep(4 * time.Second)
		messages <- "ping1"
	}()

	go func() {
		time.Sleep(7 * time.Second)
		messages <- "ping2"
	}()

	go func() {
		time.Sleep(10 * time.Second)
		messages <- "ping3"
	}()

	msg1 := <-messages
	fmt.Println("First message: ", msg1)
	msg2 := <-messages
	fmt.Println("Second message: ", msg2)

}
