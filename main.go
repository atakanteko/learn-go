package main

import (
	"fmt"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName: "Ata",
		LastName:  "Teko",
	}
	fmt.Println(user.FirstName + " " + user.LastName)
}

func saySomething(s string) (string, string) {
	fmt.Println("Function s: ", s)
	return s, "World"
}
