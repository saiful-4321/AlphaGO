package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	ToDo()
}

func ToDo() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	strArr := []string{"saiful", "Tanvir", "Raihan"}

	// append new item into array
	strArr = append(strArr, "Moazzam", "Arif", "Mostafiz")

	fmt.Println(arr, strArr)
}
