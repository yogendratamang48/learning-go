package main
import "fmt"

func main() {
	// Simple Loop
	var i int
	for i < 5 {
		fmt.Println(i)
		i++
		if i == 3 {
			continue
		}
		println("Continuing")
		
	}

	println("Second loop")
	// End condition
	for j:=1;j<5;j++ {
		println(j)
	}

}