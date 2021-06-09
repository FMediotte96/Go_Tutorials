package main

import (
	"fmt"
	"time"
)

// We've not changed anything in this function
// when compared to out previous sequential program
func compute2(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Goroutine Tutorial")

	// We've added the 'go' keyword in
	// front of both our compute functions calls
	go compute2(10)
	go compute2(10)

	var input string
	fmt.Scanln(&input)
}
