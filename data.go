package main

import (
	"log"
	"time"
)

// global variable
var s = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	Birthday    time.Time
}

func (u *User) printFirstName() string {
	return u.FirstName
}

func main() {
	log.Println("s is set to", s)

	var s2 = "six"
	log.Println("s2 is set to", s2)

	r1, r2 := saySomething("xxx")
	log.Println("saySomething func returned:", r1, r2)

	s := "eight"
	log.Println("s is now set to", s)

	user := User{
		FirstName: "Duy",
		LastName:  "Huynh",
	}

	log.Println(user)
	log.Println(user.FirstName)
	log.Println(user.printFirstName())
}

func saySomething(s string) (string, string) {
	return s, "world"
}
