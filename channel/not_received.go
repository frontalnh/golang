package main

import (
	"fmt"
	"time"
)

func main() {
	link := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		link <- "string"
	}()

	<-link
	fmt.Println("hello")

}
