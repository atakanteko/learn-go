package main

import (
	"github.com/atakanteko/learngo/helpers"
	"log"
)

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "Some Name"
	myVar.TypeNumber = 12

	log.Println(myVar)
}
