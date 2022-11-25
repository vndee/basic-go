package main

import "log"

func main() {
	for i := 0; i <= 5; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "fish", "horse", "cow", "cat"}
	for i, animal := range animals {
		log.Println(i, animal)
	}

	for _, animal := range animals {
		log.Println(animal)
	}

	anotherAnimals := make(map[string]string)
	anotherAnimals["dog"] = "Fido"
	anotherAnimals["cat"] = "Fluffy"

	for k, v := range anotherAnimals {
		log.Println(k, v)
	}

	myString := "Once upon ago"
	for i, ch := range myString {
		log.Println(i, string(ch))
	}

	type User struct {
		FirstName string
		LastName  string
	}

	var users []User
	users = append(users, User{"Duy", "Huynh"})
	users = append(users, User{"Duy", "Huynh"})
	users = append(users, User{"Duy", "Huynh"})
	users = append(users, User{"Duy", "Huynh"})

	for i, u := range users {
		log.Println(i, u)
	}
}
