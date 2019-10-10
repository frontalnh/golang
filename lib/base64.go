package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	encoded := base64.StdEncoding.EncodeToString([]byte("sdfsdfsd"))
	decoded, _ := base64.StdEncoding.DecodeString("sdfsdf")

	fmt.Println(encoded)
	fmt.Println(string(decoded))
}
