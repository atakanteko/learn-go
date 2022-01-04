package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Amy",
	}

	myVar3 := myStruct{
		FirstName: "Atakan",
	}
	log.Println("myVar is set to:", myVar.FirstName)
	log.Println("myVar2 is set to:", myVar2.FirstName)
	log.Println("myVar3 is set to:", myVar3.printFirstName())
}
