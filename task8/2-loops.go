package main 

func main() {
	var i int
	// no initializer and post condition
	for ; ; {
		if i == 5 {
			break
		}

		println(i)
		i++
	}

	println("Second loop")
	var j int
	// no initializer and post condition
	for {
		if j == 5 {
			break
		}
		println(j)
		j++
	}
}