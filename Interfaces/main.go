package main

import "fmt"

type Car interface {
	Drive()
	Stop()
}

func NewModel(arg string) Car {
	return &Lambo{arg}
}

type Lambo struct {
	Model string
}

type Bmw struct {
	Model string
}

func (l *Lambo) Drive() {
	fmt.Print("Lambo on move ")
	fmt.Println(l.Model)
}

func (B Bmw) Drive() {
	fmt.Println("BMW on move ")
	fmt.Println(B.Model)
}

func (l *Lambo) Stop() {
	fmt.Println("Stopping Lambo")
}

func main() {
	fmt.Println("Welcome to CAR interface section")

	l := Lambo{"Avantator"}
	b := Bmw{"mb34332b"}

	l.Drive()
	b.Drive()
	l.Stop()
}
