package main

import (
	"errors"
	"fmt"
)

type User struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func deleteIfNecessary(users map[string]User, name string) (deleted bool, err error) {
	if user, ok := users[name]; ok {
		if user.scheduledForDeletion {
			delete(users, name)
		}

		return user.scheduledForDeletion, nil
	} else {
		return ok, errors.New("Key not found.")
	}
}

func main() {
	myUsers := map[string]User{
		"nikilesh": User{
			name:                 "nikilesh",
			number:               27,
			scheduledForDeletion: false,
		},
		"gowramma": User{
			name:                 "GoLang",
			number:               69,
			scheduledForDeletion: true,
		},
	}

	if test, err := deleteIfNecessary(myUsers, "nikilesh"); err == nil {
		fmt.Println("Nikilesh", test)
	}

	if test, err := deleteIfNecessary(myUsers, "gowramma"); err == nil {
		fmt.Println("Gowramma", test)
	}

	fmt.Println(len(myUsers))
}
