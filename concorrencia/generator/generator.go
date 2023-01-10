package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title[^>]*>([^<]+)</title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}

func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com", "https://www.animestc.net/", "https://www.amazon.com")
	t2 := titulo("https://www.facebook.com", "https://github.com", "https://www.netflix.com", "https://www.bing.com")

	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}
