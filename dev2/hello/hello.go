package main

import (
	"fmt"
	"log"
	"yogen.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// Get Greetings message
	message, err := greetings.Hello("Yogendra")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
