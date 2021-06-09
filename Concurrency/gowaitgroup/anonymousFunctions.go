package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello World")

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		fmt.Println("Inside my goroutine")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Finished Execution")
}
