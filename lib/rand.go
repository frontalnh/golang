package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	stamp := time.Now().Unix()
	rand.Seed(int64(stamp))
	generated2 := rand.Int31()
	fmt.Println(generated2)
}
