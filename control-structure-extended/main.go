package main

import "fmt"

func main() {
	// continue and break
	flags := true
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			flags = false
			break
		}
		fmt.Println(i)
	}
	fmt.Println(flags)

	// switch case statements
	day := "sunday"

	switch day {
	case "friday":
		fmt.Println("The day is friday")
	case "sunday", "monday", "tuesday":
		fmt.Println("Boaring")
	default:
		fmt.Println("Day is another day")
	}

	switch {
	case day == "sunday":
		fmt.Println("Its fun day")
		break
	}
}
