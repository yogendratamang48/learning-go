package main

import (
	"fmt"
)


func main() {
	// We have to define struct first

	type user struct {
		ID int
		FirstName string
		LastName string
	}

	var u user
	u.FirstName = "Yogendra"
	u.LastName = "Tamang"
	u.ID = 1
	fmt.Println(u)

	// Short
	u2 := user{ ID: 2, 
		FirstName: "Arthur", 
		LastName: "Dent",
	}
	fmt.Println(u2)


}