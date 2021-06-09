package main

import "fmt"

func main() {
	// we make our anonymous function concurrent using 'go'
	go func() {
		fmt.Println("Executing my Concurrent anonymous function")
	}()

	fmt.Scanln()
}
