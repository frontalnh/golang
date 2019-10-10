package main

import (
	"fmt"
	"strconv"
)

func main() {

	number := 12312312

	stringNumber := strconv.FormatInt(int64(number), 10)

	fmt.Println(stringNumber)

	realNumber, _ := strconv.ParseInt(stringNumber, 10, 64)
	fmt.Println(realNumber)
}
