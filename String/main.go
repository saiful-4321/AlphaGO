package main

// Golang strings are muted which means string can be initialise again and again

import (
	"fmt"
	"strings"
)

func main() {
	// declaration
	var name string
	name = "Saiful Islam"

	brother := "Samim alam"

	// string concat
	fmt.Println(name + " " + brother)

	// string mutability check
	name = "Nahima khan"
	brother = "khan"

	fmt.Println(name + " " + brother)

	// checking string contains or not
	fmt.Println(strings.Contains(name, brother))

	// string replacing
	fmt.Println(strings.ReplaceAll(name, "Nahima", "Sharafat"))

	// string split
	fmt.Println(strings.Split(name, " "))
}
