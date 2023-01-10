package main

import (
	"fmt"

	"github.com/G-campos/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// Multiplexar - misturar (mensagens) num canal
func multiplexar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := multiplexar(
		html.Titulo("https://www.cod3r.com.br", "https://www.google.com"),
		html.Titulo("https://www.animestc.net/", "https://www.amazon.com"),
		// html.Titulo("https://www.facebook.com", "https://github.com"),
		// html.Titulo("https://www.netflix.com", "https://www.bing.com"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
