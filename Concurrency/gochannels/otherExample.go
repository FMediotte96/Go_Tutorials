package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 5)
	go count("sheep", c)

	var a string
	for msg := range c {
		fmt.Println(msg)
		a += msg + "|"
	}

	fmt.Println(a)
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
