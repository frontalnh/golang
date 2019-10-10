package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
}

func main() {
	person := &Person{Name: "namhoon"}

	jsonEncoded, _ := json.Marshal(person)

	fmt.Println(string(jsonEncoded))

	var dataSchema interface{}

	json.Unmarshal(jsonEncoded, &dataSchema)

	fmt.Println(dataSchema)
}
