package main 

func main() {
	// Dirty Way
	slice := []int{1, 2, 3}
	for i:=0; i <len(slice); i++ {
		println(slice[i])
	}

	for i, v := range slice {
		println(i, v)
	}
	// Looping over maps

	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for k, v := range wellKnownPorts {
		println(k, v)
	}

	// Returning only keys

	for k := range wellKnownPorts {
		println(k)
	}

	// Regurning only values
	for _, v := range wellKnownPorts {
		println(v)
	}
}