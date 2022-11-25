package main

import (
	"log"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// bad syntax
	// var myOtherMap map[string]string

	// better syntax
	myMap := make(map[string]string)

	myMap["dog"] = "Simpson"
	myMap["other-dog"] = "Cassie"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	myMap["dog"] = "Fido"
	log.Println(myMap["dog"])

	myOtherMap := make(map[string]User)

	myOtherMap["me"] = User{
		FirstName: "Duy",
		LastName:  "Huynh",
	}

	log.Println(myOtherMap["me"])

	// not recommend: user interface{} if we dont sure about the data that stored in our map
	// myMap := make(map[string]interface{})

	var mySlice []string
	mySlice = append(mySlice, "Element 1")
	mySlice = append(mySlice, "Element 2")

	log.Println(mySlice)
	log.Println(mySlice[0])

	var myOtherSlice []interface{}
	myOtherSlice = append(myOtherSlice, "String element")
	myOtherSlice = append(myOtherSlice, 1)
	log.Println(myOtherSlice)
}
