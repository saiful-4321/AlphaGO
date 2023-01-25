package main

import "fmt"

type Car struct {
	Model string
}

func main() {
	fmt.Println("Welcome to buffered channel section")

	c := make(chan int)
	c2 := make(chan *Car, 3)

	go func() {
		c2 <- &Car{"Lambo"}
		c2 <- &Car{"Buggati"}
		c2 <- &Car{"Toyota"}
		c <- 1
		c <- 2
		c <- 3
		c <- 4

		close(c)
		close(c2)
	}()

	for i := range c {
		fmt.Println(i)
	}

	for i := range c2 {
		fmt.Println(i.Model)
	}
}
