package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Items []interface{}
}

func main() {
	person := &Person{Name: "namhoon", Items: []interface{}{
		map[string]string{
			"Apple": "Applie",
		},
	}}

	fmt.Println(person)

	jsonEncoded, _ := json.Marshal(person)

	fmt.Println(string(jsonEncoded))

	var dataSchema interface{}

	json.Unmarshal(jsonEncoded, &dataSchema)

	fmt.Println(dataSchema)
}
