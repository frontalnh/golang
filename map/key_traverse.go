package main

import "fmt"

func main() {
	myMap := map[string]int{"Namhoom": 92}

	for key, value := range myMap {
		fmt.Println(key, value)
	}
}
