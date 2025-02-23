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

func sendMeArary(myArray [3]string) {
	myArray[0] = "I modified this"
}

func sendMeSlice(mySlice []string) {
	mySlice[0]  = "I modified this"
}

func printArray(array [3]string) {
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	fmt.Println()
}

func printSlice(slice []string) {
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	fmt.Println()
}

func main() {
	send("Nikilesh", 0, planPro)
	send("Nikilesh", 3, planFree)
	send("Nikilesh", 2, planPro)
	send("Nikilesh", 3, "no plan")

	myArray := [3]string {
		"This is me 1",
		"This is me 2",
		"This is me 3",
	}

	printArray(myArray)
	sendMeArary(myArray)
	printArray(myArray)

	myArray2 := myArray
	mySlice := myArray2[:]
	printSlice(mySlice)
	sendMeSlice(mySlice)
	printSlice(mySlice)
	printArray(myArray)
}
