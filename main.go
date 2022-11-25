package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	var whatToSay string
	whatToSay = "Goodbye Alice"

	fmt.Println(whatToSay)

	var i int = 1
	fmt.Println(i)

	fmt.Println(saySomething())
	st, num := saySomething()
	fmt.Println(st, num)
}

func saySomething() (string, int) {
	return "something", 42
}
