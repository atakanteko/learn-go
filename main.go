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
}
