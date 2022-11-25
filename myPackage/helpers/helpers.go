package helpers

import "log"

type SomeType struct {
	TypeName   string
	TypeNumber int
}

func Hello() {
	log.Println("Hello, this is a message from helpers package")
}
