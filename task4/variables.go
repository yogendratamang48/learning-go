package main

import (
	"fmt"
)

func main() {
	var i int
	i = 42 
	fmt.Println(i)

	// Declaring and Assigning
	var f float32 = 3.14
	fmt.Println(f)

	// Dynamic
	firstName := "Yogendra"
	fmt.Println(firstName)

	//Boolean
	b := true
	fmt.Println(b)

	// Complex Data types
	x := complex(3, 4)
	fmt.Println(x)

	r, im := real(x), imag(x)
	fmt.Println(r)
	fmt.Println(im)
}