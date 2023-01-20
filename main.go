package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Grade []int
}

func main() {
	student1 := Student{"Sadik", 29, []int{100, 90, 80}}
	student1.anything(&student1)

	fmt.Println(student1)
}

func (s *Student) anything(anything interface{}) {
	s.Name = "saif"
}
