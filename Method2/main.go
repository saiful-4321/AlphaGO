package main

import "fmt"

type Student struct {
	Name  string
	Grade []int
	Age   int
	// Average float32
	// IsAdult bool
}

func main() {
	fmt.Println("Wellcome to method learning")

	student1 := Student{"Saiful", []int{70, 78, 89, 32, 11}, 17}
	student2 := Student{"Tarikh", []int{70, 78, 89, 32, 33, 99, 90, 78, 76}, 17}

	fmt.Println(student1)

	avgGrade := student1.getAverageGrade()
	avgGrade2 := student2.getAverageGrade()

	fmt.Println(avgGrade, avgGrade2)

	student1.setAge(22)
	fmt.Println(student1)
}

func (s *Student) setAge(age int) {
	s.Age = age
}

func (s Student) getAverageGrade() float32 {
	sum := 0

	for _, val := range s.Grade {
		sum += val
	}

	return float32(sum) / float32(len(s.Grade))
}
