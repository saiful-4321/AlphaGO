package main

import "fmt"

func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	fmt.Println("Welcome to advanced interface section")
	Anything(1.22)
	Anything("Saiful")
	Anything(struct{}{})

	MyMap := make(map[interface{}]interface{})

	MyMap["Name"] = "Saiful Islam"
	MyMap["Age"] = 27
	MyMap[12] = 27

	fmt.Println(MyMap)
}
