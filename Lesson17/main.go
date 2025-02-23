package main

import (
	"fmt"
	"errors"
)

const (
	planFree = "free"
	planPro = "pro"
)

func getMessageWithRetriesForPlan(plan string) ([]string, error){
	allMessages := getMessageWithRetries()
	if plan == planPro {
		return allMessages[:], nil
	}

	if plan == planFree {
		return allMessages[0:2], nil
	}

	return nil, errors.New("Unsupported plan")
}

func getMessageWithRetries() [3]string {
	return [3]string{
		"click here to sign up",
		"pretty please click here",
		"we beg you to sign up",
	}
}

func send(name string, doneAt int, plan string) {
	fmt.Printf("Sending message to %v\n", name)

	messages, err := getMessageWithRetriesForPlan(plan)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

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
	send("Nikilesh", 0, planPro)
	send("Nikilesh", 3, planFree)
	send("Nikilesh", 2, planPro)
	send("Nikilesh", 3, "no plan")
}
