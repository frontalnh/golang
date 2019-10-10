package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	link := make(chan string)
	done := make(chan bool, 1)

	go func() {
		for i := 0; i < 10; i++ {
			link <- strconv.FormatInt(int64(i), 10)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			select {
			case <-link:
				fmt.Println("some data arrived")
			}
		}
	}()

	go func() {
		<-done
		fmt.Println("Terminated!")
	}()

	time.Sleep(5 * time.Second)
	done <- true
}
