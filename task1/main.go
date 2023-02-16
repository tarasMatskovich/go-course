package main

import (
	"fmt"
	"strings"
)

type User struct {
	name  string
	age   uint8
	email string
}

func (u User) String() string {
	return fmt.Sprintf("Name: %v\r\nAge: %v\r\nEmail: %v\r\n", u.name, u.age, u.email)
}

func main() {
	users := []User{
		{
			name:  "Mike",
			age:   32,
			email: "mike@gmail.com",
		},
		{
			name:  "John",
			age:   54,
			email: "john@gmail.com",
		},
		{
			name:  "Abobus",
			age:   2,
			email: "abobus@gmail.com",
		},
	}

	var nameToFind string
	var foundUser bool = false
	nameToFind = "john"

	for _, user := range users {
		if strings.ToLower(user.name) == strings.ToLower(nameToFind) {
			fmt.Println(fmt.Sprintf("User was found by name: %v\r\n", nameToFind))
			fmt.Println(user)
			foundUser = true
		}
	}

	if foundUser == false {
		fmt.Println(fmt.Sprintf("User was not found by name: %v\r\n", nameToFind))
	}
}
