package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {

	// This is map literal
	m1 := map[string]Vertex{
		"Bell Labs": Vertex{
			30.23, 234.123,
		},
		"Google": Vertex{
			123.312, 123.123,
		},
	}

	var m2 map[string]Vertex

	// You should initialize map before you use it!
	m2 = make(map[string]Vertex)

	m2["Bell Labs"] = Vertex{123.123, 123.123}

	fmt.Println(m1["Bell Labs"])
	fmt.Println(m2["Bell Labs"])

	delete(m2, "Bell Labs")

	// ok tells us that that value is exists
	v, ok := m2["Bell Labs"]
	fmt.Println(v, ok)
	fmt.Println(m2["Bell Labs"])
}
