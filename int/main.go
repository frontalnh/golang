package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := 123123

	fmt.Println(strconv.FormatInt(int64(number), 10))
}
