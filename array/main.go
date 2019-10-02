package main

import "fmt"

func main() {
	arr := []int{2, 3, 4, 5, 6}

	fmt.Println("arr: ", arr)

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for i, v := range arr {
		fmt.Printf("%dth element is %d \n", i, v)
	}
}
