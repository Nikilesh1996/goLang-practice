package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Manager struct {
	// This is an embedded structure
	Person
	Level int
}

// This is an embedded structure

func main() {
	// These are anonymous structures

	myPerson := struct {
		FirstName string
		LastName  string
		Age       int
	}{
		FirstName: "Nikilesh",
		LastName:  "Raj",
		Age:       28,
	}

	fmt.Println(myPerson.FirstName)

	// Embedded structure example
	student := Person{
		FirstName: "Nkilesh",
		LastName:  "Raj",
		Age:       28,
	}

	// this is how you initialize embedded struct
	hisManager := Manager{
		Person: Person{
			FirstName: "Gowri",
			LastName:  "Baalambal",
			Age:       27,
		},
		Level: 1,
	}

	fmt.Println(student.FirstName)
	fmt.Println(hisManager.FirstName)
}
