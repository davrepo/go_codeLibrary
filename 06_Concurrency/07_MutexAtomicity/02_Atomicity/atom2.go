package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync/atomic"

	"golang.org/x/sync/errgroup"
)

func downloadsSize(urls []string) (int, error) {
	var dlSize int64
	// errgroup.Group is a convenient way to parallelize independent tasks
	// to know when all tasks are done and to collect their errors
	var g errgroup.Group

	for _, url := range urls {
		// construct a new variable to avoid closure capture of the loop variable
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			log.Print(url)
			size, err := downloadSize(url)
			if err != nil {
				return err
			}

			atomic.AddInt64(&dlSize, int64(size))
			return nil
		})
	}

	// Wait blocks until all function calls from the Go method have returned,
	// then returns the first non-nil error (if any) from them.
	if err := g.Wait(); err != nil {
		return 0, err
	}

	return int(dlSize), nil
}

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
