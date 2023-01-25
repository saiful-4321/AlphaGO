package main

import (
	"fmt"
	"time"
)

// The pinger print a ping and wait for pong
func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		<-pinger
		fmt.Println("Ping")
		time.Sleep(time.Second)
		ponger <- 1
	}
}

func ponger(pinger chan<- int, ponger <-chan int) {
	for {
		<-ponger
		fmt.Println("Pong")
		time.Sleep(time.Second)
		pinger <- 1
	}
}

func main() {
	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)

	ping <- 1
	select {}
}
