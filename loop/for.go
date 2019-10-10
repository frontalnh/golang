package main

import "fmt"

func main() {
	arr := []string{"first", "second"}

	for i, v := range arr {
		fmt.Printf("%dth element is %s", i, v)
	}

	n := 1
	for n < 5 {
		fmt.Println(n)
		n++
	}
}
