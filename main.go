package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
[
    {
        "first_name": "Clark",
        "last_name": "Kent",
        "hair_color": "black",
        "has_dog": true
    },
    {
        "first_name": "Bruce",
        "last_name": "Wayne",
        "hair_color": "black",
        "has_dog": false
    }
]`
	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}
	log.Printf("unmarshalled: %v", unmarshalled)

	//write json from a struct
	var mySlice []Person

	var marshall1 Person
	marshall1.FirstName = "Carla"
	marshall1.LastName = "Dayton"
	marshall1.HairColor = "Brown"
	marshall1.HasDog = false

	mySlice = append(mySlice, marshall1)

	var marshall2 Person
	marshall2.FirstName = "James"
	marshall2.LastName = "Fusco"
	marshall2.HairColor = "Black"
	marshall2.HasDog = true

	mySlice = append(mySlice, marshall2)

	for _, item := range mySlice {
		log.Println(item)
	}
	newJSon, err := json.MarshalIndent(mySlice, "", "    ")
	if err != nil {
		log.Println("error marshalling", err)
	}
	fmt.Print(string(newJSon))
}
