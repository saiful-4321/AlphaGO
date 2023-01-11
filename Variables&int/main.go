package main

import "fmt"

func main() {
	fmt.Println("Welcome to variable section")

	var m1 int //single variable declaration

	// multiple varible declaration
	var (
		m2 int32
		m3 int64
	)

	// type casting int32(value)
	fmt.Println(int32(m1) + m2 + int32(m3))

	// initialise variable
	a1 := 22
	a2 := 44

	fmt.Println(a1 + a2)

}
