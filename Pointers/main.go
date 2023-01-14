package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers section.")
	// name := "saiful"
	// ptr := &name
	// ptr2 := &name
	// fmt.Println(ptr)
	// fmt.Println(*ptr2)

	name, greetings := "saiful", "Hello"

	fmt.Println(name, greetings)
	swap(&name, &greetings)
	fmt.Println(name, greetings)
}

func swap(name, greetings *string) {
	var temp string

	temp = *greetings
	*greetings = *name
	*name = temp
}
