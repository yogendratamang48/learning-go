package main

import "fmt"
import "rsc.io/quote"


// Structs - composite
type info struct {
	result string
}

func infoRepo(a int) (string, error) {
	return "This is the intro package", nil
}

func main() {
	var introMsg string = "Hello from Go Toolchain"
	fmt.Println(introMsg)
	// String Formatting
	fmt.Println("Go Reports: %s\n ", introMsg)
	// String Formatting
	fmt.Println("[Avoid] Go Reports: " + introMsg)
	mssg, err := infoRepo(2)

	if err != nil {
		fmt.Println("Go reports error condition: %v\n ", err)
	}
	fmt.Println("Go reports error condition: %v\n ", mssg)

	// Go Arrays
	xs := []float64 {98, 100, 87}
	total := 0.0
	for _, v:= range xs {
		total += v
	}
	fmt.Println(total)

	// Go Structs composite type
	var s = info{}

	// Generate a pointer to s
	sp := &s
	sp.result = "Set a struct pointer value"
	fmt.Println("Go Reports: %+v\n", sp.result )
	// Using package
	fmt.Println(quote.Go())

}
