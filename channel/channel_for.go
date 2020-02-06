package main

import "fmt"

func main() {
	messages := make(chan int, 3)

	messages <- 1
	messages <- 1
	messages <- 1

	for o := range messages {
		fmt.Println(o)
	}
}
