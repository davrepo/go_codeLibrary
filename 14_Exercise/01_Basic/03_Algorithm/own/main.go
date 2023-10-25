package main

import (
	"flag"
	"log"
)

// coin contains the name and value of a coin
type coin struct {
	name  string
	value float64
}

// coins is the list of values available for paying.
var coins = []coin{
	{name: "1 pound", value: 1},
	{name: "50 pence", value: 0.50},
	{name: "20 pence", value: 0.20},
	{name: "10 pence", value: 0.10},
	{name: "5 pence", value: 0.05},
	{name: "1 penny", value: 0.01},
}

// calculateChange returns the minimum number of coins required to pay the given amount.
func calculateChange(amount float64) map[coin]int {
	// divide amount by successive biggest coin value, return remainder
	// insert coin name and count into map
	// repeat with remainder until remainder is 0

	// create map to hold change
	change := make(map[coin]int)

	// loop over coins
	for _, c := range coins {
		if amount >= c.value {
			// calculate number of coins of this value
			count := int(amount / c.value)
			// calculate remainder
			amount = amount - float64(count)*c.value
			// add coin to change map
			change[c] = count
		}
	}

	return change
}

// printCoins prints all the coins in the slice to the terminal.
func printCoins(change map[coin]int) {
	if len(change) == 0 {
		log.Println("No change found.")
		return
	}
	log.Println("Change has been calculated.")
	for coin, count := range change {
		log.Printf("%d x %s \n", count, coin.name)
	}
}

func main() {
	amount := flag.Float64("amount", 0.0, "The amount you want to make change for")
	flag.Parse()
	change := calculateChange(*amount)
	printCoins(change)
}
