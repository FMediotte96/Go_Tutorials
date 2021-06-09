package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValue2(c chan int) {
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value: {}", value)
	time.Sleep(1000 * time.Millisecond)
	c <- value
	fmt.Println("Only Executes after another goroutine performs a receive on the channel")
}

func main() {
	fmt.Println("Go Channel Tuttorial")

	valueChannel := make(chan int)
	defer close(valueChannel)

	go CalculateValue2(valueChannel)
	go CalculateValue2(valueChannel)

	values := <-valueChannel
	fmt.Println(values)
}
