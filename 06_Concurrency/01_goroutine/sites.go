// Get content type of sites
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)
}

// sequential version
func siteSequential(urls []string) {
	for _, url := range urls {
		returnType(url)
	}
}

// concurrent version
func sitesConcurrent(urls []string) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			returnType(url)
			wg.Done()
		}(url)
	}
	wg.Wait() // (1)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/ip",
	}

	start := time.Now()
	siteSequential(urls)
	fmt.Println(time.Since(start))

	start = time.Now()
	sitesConcurrent(urls)
	fmt.Println(time.Since(start))
}
