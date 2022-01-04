package main

import (
	"github.com/atakanteko/learngo/helpers"
	"log"
)

const numPool = 10

func CalculateValue(intChannel chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChannel <- randomNumber
}
func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "Some Name"
	myVar.TypeNumber = 12

	log.Println(myVar)

	intChannel := make(chan int)
	defer close(intChannel)

	go CalculateValue(intChannel)
	rndNumber := <-intChannel
	log.Println(rndNumber)
}
