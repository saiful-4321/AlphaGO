package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// fmt.Println("Welcome to select block section")
	start := make(chan int)
	quit := make(chan struct{})

	go func() {
		start <- 1
		quit <- struct{}{}
	}()

	go Select(start, quit)
	select {}
}

func Select(start chan int, quit chan struct{}) {
	// select case are like switch case
	// select used for channel and it is asyncronous
	for {
		time.Sleep(time.Second)
		select {
		case <-start:
			fmt.Println("Start")
		case <-quit:
			fmt.Println("Quit")
			os.Exit(0)
		}
	}
}
