package main

import "fmt"

func main() {
	// initialize variables here
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	smsSendingLimit = 0
	costPerSMS = 0
	hasPermission = false
	username = ""

	// This is string concatenation
	var firstString string = "Hello"
	var secondString string = "World"

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
	fmt.Printf("%q", firstString+" "+secondString)
}
