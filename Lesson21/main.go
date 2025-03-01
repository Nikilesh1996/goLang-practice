package main

import (
	"errors"
	"fmt"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]User, error) {
	userMap := make(map[string]User)

	if len(names) != len(phoneNumbers) {
		return nil, errors.New("Invalid sizes")
	}

	for index, name := range names {
		userMap[name] = User{
			name:        name,
			phoneNumber: phoneNumbers[index],
		}
	}

	return userMap, nil
}

type User struct {
	name        string
	phoneNumber int
}

// Lesson for maps,
// Map is key value pair.
func main() {
	userMaps, err := getUserMap(
		[]string{"Nikilesh", "Niveda", "Gowri"},
		[]int{12323423423, 1234234234, 12312312},
	)

	if err == nil {
		fmt.Println("These are the users and their phoneNumbers")
		for key, value := range userMaps {
			fmt.Printf("Key: %v; Value: %v;\n", key, value)
		}
	} else {
		fmt.Println(err.Error())
	}
}
