package main

import (
	"fmt"
	"sync"
)

func main() {
	// wait groups
	wg := &sync.WaitGroup{}
	// lemda or annonimous function
	// add done and wait functionality
	wg.Add(2)
	go func() {
		fmt.Println("hello from lemda function")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello from another lemda function")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("wellcome to concurrency extended section")
}
