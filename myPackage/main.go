package main

import (
	"log"

	"github.com/vndee/mypackage/helpers"
)

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "Some Type"
	log.Println(myVar)

	helpers.Hello()
}
