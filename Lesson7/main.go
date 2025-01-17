package main

import "fmt"

// This is to demostrate the usage of structure
type MessageToSend struct {
	message string

	// This is a nested structure
	sender   User
	receiver User
}

type User struct {
	name   string
	number int
}

func main() {
	message := MessageToSend{
		message: "Hello world!",
		sender: User{
			name:   "Nikilesh",
			number: 9,
		},
		receiver: User{
			name:   "Niveda",
			number: 7,
		},
	}

	fmt.Println(message)
}
