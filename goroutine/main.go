package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to goroutine section")
	go heavy()
	go superHeavy()
	fmt.Println("Gotcha")
	// time.Sleep(time.Second * 5)
	select {}
}

func heavy() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("called heavy routine")
	}
}

func superHeavy() {
	for {
		time.Sleep(time.Second * 3)
		fmt.Println("Called super heavy routine")
	}
}
