package main

import (
	"fmt"
)

func getMessageWithRetries() [3]string {
	return [3]string{
		"click here to sign up",
		"pretty please click here",
		"we beg you to sign up",
	}
}

func send(name string, doneAt int) {
	fmt.Printf("Sending message to %v\n", name)

	messages := getMessageWithRetries()
	for i := 0; i < len(messages); i++ {
		fmt.Printf("sending: %v\n", messages[i])
		if i == doneAt {
			break
		}

		if i == len(messages)-1 {
			fmt.Println("Complete failure")
		}
	}
}

func main() {
	send("Nikilesh", 0)
	send("Nikilesh", 1)
	send("Nikilesh", 2)
	send("Nikilesh", 3)
}
