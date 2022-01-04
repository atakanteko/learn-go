package main

import (
	"log"
)

func main() {
	var mySlice []string

	mySlice = append(mySlice, "Kemal")
	mySlice = append(mySlice, "Melih")
	mySlice = append(mySlice, "Ahmet")

	log.Println(mySlice)

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	log.Println(numbers)
	log.Println(numbers[4:7])

	animals := []string{"dog", "fish", "horse", "cow", "cat", "duck"}

	for _, animal := range animals {
		log.Println(animal)
	}

	soccerTeams := make(map[string]string)

	soccerTeams["team1"] = "Galatasaray"
	soccerTeams["team2"] = "Fenerbahçe"
	soccerTeams["team3"] = "Beşiktaş"

	for teamName, item := range soccerTeams {
		log.Println(teamName, item)
	}

	type PersonelInfo struct {
		FirstName string
		LastName  string
		Age       int
		Email     string
	}

	var users []PersonelInfo
	users = append(users, PersonelInfo{"Elif", "Aksu", 22, "elifaksu@gmail.com"})
	users = append(users, PersonelInfo{"Kemal", "Ekin", 22, "kemalekin@gmail.com"})
	users = append(users, PersonelInfo{"Nedim", "Gazi", 22, "nedimgazi@gmail.com"})
	users = append(users, PersonelInfo{"Börte", "Kahraman", 22, "bortekahraman@gmail.com"})

	for _, item := range users {
		log.Println(item.FirstName, " ", item.LastName, " ", item.Age, " ", item.Email)
	}
}
