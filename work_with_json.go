package main

import (
	"json"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	myJson := `
	[
		{
			"name": "Alice",
			"age", 2
		},
		{
			"name": "Bob",
			"age": 3
		}
	]
	`

	var unmarshalled []Person
	err := json.Unmarshalled([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshaled: %v", unmarshalled)

	// write json from a struct
	var mySlice []Person

	var m1 Person
	m1.Name = "Wally"
	m1.Age = 12

	mySlice = append(mySlice, m1)

	newJson, err := json.MarshalIndent(mySlice, "", "    ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	log.Println(string(newJson))
}
