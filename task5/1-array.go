package main

import (
	"fmt"
)

func main() {
	// Simple Declaration
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	arr2 := [3]int{4, 5, 6}
	fmt.Println(arr2)

}