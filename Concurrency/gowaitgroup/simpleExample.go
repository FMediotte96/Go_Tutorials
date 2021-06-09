package main

import (
	"fmt"
	"sync"
)

func myFunc(wg *sync.WaitGroup) {
	fmt.Println("Inside my goroutine")
	wg.Done()
}

func main() {
	fmt.Println("Hello world")

	var wg sync.WaitGroup
	wg.Add(1)
	go myFunc(&wg)
	wg.Wait() //block the execution util the goroutines in the waitgroup have successfully completed

	fmt.Println("Finished Execution")
}
