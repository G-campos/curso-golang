package main

import (
	"fmt"
	"time"
)

// Channel (canal) - é a forma de comunicação entre goroutines
// channel é um tipo

func doisTresQuatroVezes(base int, c chan int) {
	fmt.Println("X")
	time.Sleep(time.Second)
	c <- 2 * base //enviando dados para o cana

	fmt.Println("Y")
	time.Sleep(time.Second)
	c <- 3 * base

	fmt.Println("Z")
	time.Sleep(time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)
	fmt.Println("A")

	a, b := <-c, <-c //recebendo dados do canal
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<-c)
}
