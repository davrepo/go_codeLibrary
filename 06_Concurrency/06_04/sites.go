// Get content type of sites (using channels)
package main

import (
	"fmt"
	"net/http"
)

func returnType(url string, channel chan string) {
	resp, err := http.Get(url)
	if err != nil {
		channel <- fmt.Sprintf("%s -> error: %s", url, err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	channel <- fmt.Sprintf("%s -> %s", url, ctype)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	// Create response channel
	ch := make(chan string)
	for _, url := range urls {
		go returnType(url, ch)
	}

	for range urls { // same as for _, _ := range urls
		out := <-ch
		fmt.Println(out)
	}
}
