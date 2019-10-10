package main

import "fmt"

func main() {
	myMap := map[string]string{}

	v, found := myMap["hello"]

	fmt.Println(v, found)
}
