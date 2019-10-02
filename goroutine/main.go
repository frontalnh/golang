package main

import "fmt"

func asyncStart() error {
	go func(word string) {
		fmt.Println(word)
	}("hello")

	return nil
}

func main() {
	asyncStart()
}
