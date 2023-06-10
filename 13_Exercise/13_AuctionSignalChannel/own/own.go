package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// Bidding for resources
// Given a list of recourses,
// simulate concurrent management and allocation of bids for recourses.
// assumptions:
// - all bidders start with the same amount of money
// - all bidders bid on all items
// - all bids in integer amount
// - if bidder is out of money, they will place no more bids
// - no minimum prices for items
// - items are sold one by one sequentially

const bidderCount = 10   // number of bidders
const walletAmount = 250 // initial money for all bidders

// items to be bid (slice of strings)
var items = []string{
	"The \"Best Gopher\" trophy",
	"The \"Learn Go with Adelina\" experience",
	"Two tickets to a Go conference",
	"Signed copy of \"Beautiful Go code\"",
	"Vintage Gopher plushie",
}

// bid struct pairs the bidder id and the amount they want to bid
type bid struct {
	bidderID string
	amount   int
}

// auctioneer receives bids and announces winners
type auctioneer struct {
	bidders map[string]*bidder
	bids    chan *bid // buffered channel to receive bids, size of bidderCount
}

// runAuction and manages the auction for all the items to be sold
func (a *auctioneer) runAuction(wg *sync.WaitGroup) {
	for _, item := range items {
		log.Printf("Opening bids for %s!\n", item)

		wg.Add(bidderCount)

		// create a channel to receive bids for this item
		for _, b := range a.bidders {
			// send a bid to the channel
			// each bidder places a bid concurrently
			go b.placeBid(item, a.bids, wg) // bid sent to a.bids channel
		}

		// wait for all bidders to place their bids
		wg.Wait()

		highestBid := &bid{}
		for i := 0; i < len(a.bidders); i++ {
			b := <-a.bids // receive a bid from the channel
			if b.amount > highestBid.amount {
				highestBid = b
			}
		}

		if highestBid.bidderID != "" {
			// get the bidder struct from map, and assign it to winner
			winner := a.bidders[highestBid.bidderID]
			winner.payBid(highestBid.amount)
			log.Printf("The winner of %s is %s with a bid of %d!\n", item, winner.id, highestBid.amount)
		}
	}
}

// bidder is a type that holds the bidder id and wallet
type bidder struct {
	id     string
	wallet int
}

// placeBid generates a random amount and places it on the bids channels
// Change the signature of this function as required
func (b *bidder) placeBid(item string, bids chan *bid, wg *sync.WaitGroup) {
	defer wg.Done() // decrement the wait group counter

	if b.wallet <= 0 {
		return
	}

	amount := getRandomAmount(b.wallet)
	bids <- &bid{ // make a bid struct, and send it to the channel
		bidderID: b.id,
		amount:   amount,
	}
}

// getRandomAmount generates a random integer amount up to max
func getRandomAmount(max int) int {
	return rand.Intn(int(max)) + 1
}

// subtracts the bid amount from the wallet of the auction winner
func (b *bidder) payBid(amount int) {
	b.wallet -= amount
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.Println("Welcome to auction.")

	// create a map of bidders
	bidders := make(map[string]*bidder, bidderCount)
	for i := 0; i < bidderCount; i++ {
		id := fmt.Sprintf("Bidder %d", i)
		bidders[id] = &bidder{
			id:     id,
			wallet: walletAmount,
		}
	}

	// create an auctioneer
	a := auctioneer{
		bidders: bidders,
		bids:    make(chan *bid, bidderCount), // buffered channel
	}

	var wg sync.WaitGroup
	a.runAuction(&wg)

	log.Println("The auction has finished!")
}
