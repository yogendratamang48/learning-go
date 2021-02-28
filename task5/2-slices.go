package main

import (
	"fmt"
)

func main() {
	// Simple Declaration
	arr := [3]int{4, 5, 6}

	slice := arr[:]
	arr[1] = 42
	slice[2] = 27
	fmt.Println(arr, slice)
	// Slice Kind of pointer to Array

	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)
	slice2 = append(slice2, 4, 5, 6, 7)
	fmt.Println(slice2)

	// Like python slices
	s2 := slice2[1:]
	s3 := slice2[:2]
	s4 := slice2[1:2]
	fmt.Println(s2, s3, s4)

}