package main

import (
	"fmt"
	"yogen.com/greetings"
)

func main() {
	// Get Greetings message
	message := greetings.Hello("Yogendra")
	fmt.Println(message)
}
