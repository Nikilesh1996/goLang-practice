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

func computeIfUserIsValid(messageToSend MessageToSend) bool {
	if messageToSend.sender.name == "" {
		return false;
	}

	if messageToSend.receiver.name == "" {
		return false;
	}

	if messageToSend.sender.number == 0 {
		return false;
	}

	if messageToSend.receiver.number == 0 {
		return false;
	}

	return true;
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
	fmt.Println(computeIfUserIsValid(message))
}
