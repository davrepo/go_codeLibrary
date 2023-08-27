package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func downloadSize(url string) (int, error) {
	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("bad status: %d %s", resp.StatusCode, resp.Status)
	}

	return strconv.Atoi(resp.Header.Get("Content-Length"))
}

func downloadsSize(urls []string) (int, error) {
	var total int
	var wg sync.WaitGroup
	wg.Add(len(urls))    // total number of jobs
	ch := make(chan int) // channel to receive size of each download

	for i := 0; i < len(urls); i++ {
		go func(url string) {
			defer wg.Done()
			size, err := downloadSize(url)
			if err != nil {
				log.Println(err)
				return
			}
			ch <- size
		}(urls[i])
	}

	for size := range ch {
		total += size
	}

	wg.Wait()
	close(ch)

	return total, nil
}

func gen2020URLs() []string {
	var urls []string
	urlTemplate := "https://s3.amazonaws.com/nyc-tlc/trip+data/%s_tripdata_2020-%02d.csv"
	for _, vendor := range []string{"yellow", "green"} {
		for month := 1; month <= 12; month++ {
			url := fmt.Sprintf(urlTemplate, vendor, month)
			urls = append(urls, url)
		}
	}
	return urls
}

func main() {
	urls := gen2020URLs()
	size, err := downloadsSize(urls)
	if err != nil {
		log.Fatal(err)
	}

	sizeGB := float64(size) / (1 << 30)
	fmt.Printf("size = %.2fGB\n", sizeGB)
}
