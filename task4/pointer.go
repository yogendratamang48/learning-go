package main

import (
	"fmt"
)

func pointer1() {
	var firstName *string = new(string)
	*firstName = "Yogendra" 
	fmt.Println(firstName, *firstName)


}


func main() {
	firstName := "Yogendra"
	fmt.Println(firstName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Hari"
	fmt.Println(firstName, ptr, *ptr)
}
	
	