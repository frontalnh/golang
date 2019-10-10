package main

import "fmt"

func main() {
	arr := []int{2, 3, 4, 5, 6}

	arr = append(arr, 7)

	fmt.Println("arr: ", arr)
	fmt.Println(arr[0:3])
	fmt.Println(arr[0:len(arr)])
	fmt.Println(arr[0:])

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for i, v := range arr {
		fmt.Printf("%dth element is %d \n", i, v)
	}
}
