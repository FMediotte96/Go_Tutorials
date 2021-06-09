package main

import (
	"fmt"
)

func main() {
	a := make(chan int, 1)   //bufferedChannel
	b := make(chan struct{}) //para determinar cuando termina mi proceso

	go escuchar2(a, b)
	publicar2(a)

	b <- struct{}{}

}

//<-chan con esto indicamos que el canal solo va a recibir
func escuchar2(a <-chan int, b <-chan struct{}) {
	for {
		select {
		case i := <-a:
			fmt.Println(i)
		case <-b:
			break
		}
	}
}

//chan<- con esto indicamos que el canal solo va a enviar datos al canal
func publicar2(a chan<- int) {
	a <- 7
	a <- 8
	a <- 4
	a <- 6
	a <- 1
	a <- 61
	a <- 12
	a <- 81
}
