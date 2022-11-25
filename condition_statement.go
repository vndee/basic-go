package main

import "log"

func main() {
	isTrue := true
	if isTrue {
		log.Println("True")
	} else {
		log.Println("False")
	}

	myVar := 1
	switch myVar {
	case 0:
		log.Println("it is zero")

	case 1:
		log.Println("it is one")

	case 2:
		log.Println("it is two")

	default:
		log.Println("it is default")
	}
}
