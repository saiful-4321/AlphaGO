package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to control structures section")

	f := true
	flag := &f

	// if else statement
	if flag == nil {
		fmt.Println("Flag is nil")
	} else if *flag {
		fmt.Println("Flag is true")
	} else {
		fmt.Println("Flag is false")
	}

	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println("The number is", i)
	}

	// using range
	students := []string{"saiful", "ratikh", "asabul"}
	for i, val := range students {
		fmt.Println(i, val)
	}

	// using map
	mymap := make(map[string]interface{})
	mymap["name"] = "saiful"
	mymap["designation"] = "software engineer"
	mymap["org"] = "Paywll"

	for key, value := range mymap {
		fmt.Printf("Key is %s value is %v ", key, value)
	}
}
