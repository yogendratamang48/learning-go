package main

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	u1 := User{
		ID:        1,
		FirstName: "Yogendra",
		LastName:  "Tamang",
	}
	u2 := User{
		ID:        2,
		FirstName: "Hari",
		LastName:  "Thapa",
	}

	if u1.ID == u2.ID {
		println("Same User !")
	} else {
		println("Different Users!")
	}
}
