package main

import "fmt"

// This is to demostrate the usage of structure
type messageToSend struct {
	message string

	// This is a nested structure
	sender   user
	receiver user
}

type user struct {
	name   string
	number int
}

func main() {
	message := messageToSend{
		message: "Hello world!",
		sender: user{
			name:   "Nikilesh",
			number: 9,
		},
		receiver: user{
			name:   "Niveda",
			number: 7,
		},
	}

	fmt.Println(message)
}
