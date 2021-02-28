package main
import (
	"fmt"
)

const (
	first = iota + 6
	second = iota << 3

)
func main() {
	fmt.Println(first, second)
}