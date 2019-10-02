package main

import (
	"fmt"
	"math"
)

type Person struct {
	name string
}

func (p *Person) Tell(word string) {
	fmt.Println(word)
}

func (p *Person) Sqrt(value float64) float64 {

	return math.Sqrt(value)
}

func main() {
	person := &Person{name: "namhoon"}

	person.Tell("Hello")

	fmt.Println(person)
	fmt.Println(*person)
}
