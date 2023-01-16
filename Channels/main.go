package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to channel section")
	// var channel chan int
	channel := make(chan int)

	// send in a goroutine
	go func() {
		channel <- 1
	}()

	val := <-channel
	fmt.Println(val)

	// just incase
	time.Sleep(time.Second * 2)

	go func() {
		channel <- 2
	}()

	val = <-channel
	fmt.Println(val)

	fmt.Println(channel)
}
