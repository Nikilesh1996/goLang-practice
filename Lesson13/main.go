package main

import (
	"fmt"
)

type User struct {
	userId, userFirstName, userLastName string
}

type UserProfile struct {
	userAddress string
}

func getUser() (User, error) {
	tempUser := User{
		userId:        "123",
		userFirstName: "Nikilesh",
		userLastName:  "Raj N",
	}

	return tempUser, nil
}

func getUserProfile() (userProfile UserProfile, e error) {
	tempUserProfile := UserProfile{
		userAddress: "Bengaluru",
	}

	return tempUserProfile, nil
}

func sendSms(message string) (float64, error) {
	const maxTextLength = 25
	const costPerCharacter = 0.0002

	if len(message) > maxTextLength {
		return 0.0, fmt.Errorf("Text length exceeds max text length")
	}

	return (float64(len(message)) * costPerCharacter), nil

}

func sendSmsToCouple(messageToCustomer string, messageToSpouse string) (float64, error) {
	costForCustomer, err := sendSms(messageToCustomer)
	if err != nil {
		return 0.0, err
	}

	costForSpouse, err := sendSms(messageToSpouse)
	if err != nil {
		return 0.0, err
	}

	return costForCustomer + costForSpouse, nil
}

func main() {
	// How does go handle errors

	user, err := getUser()
	if err != nil {
		fmt.Println("User not found")
	}

	userProfile, err := getUserProfile()
	if err != nil {
		fmt.Println("User profile not found")
	}

	fmt.Printf("%s %s \n", user.userFirstName, userProfile.userAddress)

	cost, err := sendSmsToCouple("This is for the customer", "This is for spouse")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("The cost for the message is", cost)
}
