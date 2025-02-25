// Context example
package main

import (
	"context"
	"fmt"
	"time"
)

type Bid struct {
	AdURL string
	Price float64
}

func bestBid(url string) Bid {
	time.Sleep(20 * time.Millisecond)

	return Bid{
		AdURL: "https://adsЯus.com/19",
		Price: 0.05,
	}
}

var defaultBid = Bid{
	AdURL: "https://adsЯus.com/default",
	Price: 0.02,
}

func findBid(ctx context.Context, url string) Bid {
	ch := make(chan Bid, 1) // buffered to avoid goroutine leak
	go func() {
		ch <- bestBid(url)
	}()

	select {
	// in case we got a bid from the channel
	// it means the algorithm finished at the allocated time
	case bid := <-ch:
		return bid
	// in case we didn't get a bid from the channel
	// it means the algorithm didn't finish at the allocated time
	// so we return the default bid
	case <-ctx.Done():
		return defaultBid
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	url := "https://http.cat/418"
	bid := findBid(ctx, url)
	fmt.Println(bid)

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	url = "https://http.cat/404"
	bid = findBid(ctx, url)
	fmt.Println(bid)
}
