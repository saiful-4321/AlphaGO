package main

import (
	"fmt"
)

type Details struct {
	Name       string
	Class      string
	Profession string
}

func main() {
	fmt.Println("Welcome back again")

	var teach Details

	teach.updateDetails("Saiful", "Eleven", "Teaching")
}

func (teacher Details) updateDetails(name string, class string, profession string) {
	teacher.Name = name
	teacher.Class = class
	teacher.Profession = profession

	fmt.Println(teacher)
}
