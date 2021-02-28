package main

import (
	"github.com/yogendratamang48/webservice/models"
	"fmt"
)

func main() {
	u := models.User{
		ID: 2,
		FirstName: "Yogendra",
		LastName: "Tamang",
	}
	fmt.Println(u)
}