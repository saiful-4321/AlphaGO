package main

import "fmt"

type Student struct {
	Name      string
	Age       int
	isMarried bool
	Grade     []int
}

func main() {
	fmt.Println("Welcome to struct section")

	student := Student{"Saiful", 19, false, []int{100, 39, 79, 89, 69, 98}}
	// student = Student{"Samim", 26, true, []int{100, 39, 79, 89, 61, 98}}
	// student = Student{"Nusaifa", 2, false, []int{10, 9, 8, 7, 6, 5, 4, 3}}

	student.Print()
	student.Learning()
	fmt.Println(student.GetName())
}

func (st Student) Print() {
	fmt.Println(st)
}

func (st Student) Learning() {
	fmt.Println("Learning golang")
}

func (st Student) GetName() string {
	return st.Name
}
