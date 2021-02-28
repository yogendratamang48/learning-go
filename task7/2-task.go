package main 

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to function")
	port := 3000
	_, err := startWebServer(port, 2)
	fmt.Println(err)
}

func startWebServer(port int, numberOfRetries int) (int, error) {
	fmt.Println("Starting webserver...")
	fmt.Println("Server Started", port)
	fmt.Println("Retries", numberOfRetries)
	return port, nil
}