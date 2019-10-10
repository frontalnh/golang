package main

import "fmt"

type Person struct {
	tasks chan string
}

func main() {
	person := &Person{make(chan string)}

	go func() {
		person.tasks <- "task1"
	}()

	task := <-person.tasks

	fmt.Println(task)
}
