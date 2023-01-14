package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers section.")

	name := "saiful"

	ptr := &name
	fmt.Println(ptr)

	name = "Islam"
	ptr = &name

	fmt.Println(name)
	fmt.Println(ptr)
}
