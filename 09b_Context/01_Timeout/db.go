// Dummy bid model
package main

import "time"

var state = 0

func bestBid(url string) Bid {
	state = 1 - state // toggle state
	if state == 1 {
		// fast bid
		time.Sleep(2 * time.Millisecond)
		return Bid{
			Price: 0.035,
			URL:   "https://j.mp/3f3Dpkb",
		}
	}
	// bidTimeout is 10ms
	// so this will always timeout
	time.Sleep(bidTimeout + 20*time.Millisecond)
	// so this block of code will never be executed
	return Bid{
		Price: 0.018,
		URL:   "https://j.mp/39oEJe7",
	}
}
