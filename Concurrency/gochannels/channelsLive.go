package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make(chan int, 1) //bufferedChannel

	wg := sync.WaitGroup{}
	wg.Add(2)

	go escuchar(a, &wg)
	go publicar(a, &wg)

	wg.Wait()
}

//<-chan con esto indicamos que el canal solo va a recibir
func escuchar(a <-chan int, wg *sync.WaitGroup) {
	for i := range a {
		fmt.Println(i)
	}
	wg.Done()
}

//chan<- con esto indicamos que el canal solo va a enviar datos al canal
func publicar(a chan<- int, wg *sync.WaitGroup) {
	a <- 7
	a <- 8
	a <- 4
	a <- 6
	a <- 1
	a <- 61
	a <- 12
	a <- 81
	close(a)
	wg.Done()
}
