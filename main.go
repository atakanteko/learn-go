package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myVar := make(map[string]string)

	myVar["dogName"] = "Duman"
	myVar["myFavoriteMovie"] = "Harry Potter"

	log.Println(myVar["dogName"])
	log.Println(myVar["myFavoriteMovie"])

	myAnotherVar := make(map[string]int)

	myAnotherVar["smallestPrimeNumber"] = 2

	log.Println(myAnotherVar["smallestPrimeNumber"])

	myStructMap := make(map[string]User)

	personelInfo := User{
		FirstName: "Atakan",
		LastName:  "Tekoğlu",
	}
	myStructMap["personelInfo"] = personelInfo

	log.Println(myStructMap["personelInfo"].FirstName + " " + myStructMap["personelInfo"].LastName)
}
