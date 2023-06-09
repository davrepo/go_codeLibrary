package main

import (
	"flag"
	"log"
	"time"
)

var expectedFormat = "2006-01-02"

// time.Parse()
// time.Now()
// time.Until()

// run code:
// go run . -bday 2023-05-01
// Error parsing date: <nil>
// go run . -bday 2024-05-01
// You have 326 days until your birthday. Hurray!

// parseTime validates and parses a given date string.
func parseTime(target string) time.Time {
	pt, err := time.Parse(expectedFormat, target)
	if err != nil || time.Now().After(pt) {
		log.Fatalf("Error parsing date: %v", err)
	}

	return pt
}

// calcSleeps returns the number of sleeps until the target.
func calcSleeps(target time.Time) float64 {
	return time.Until(target).Hours() / 24
}

func main() {
	bday := flag.String("bday", "", "Your next bday in YYYY-MM-DD format")
	flag.Parse()
	target := parseTime(*bday)
	log.Printf("You have %d days until your birthday. Hurray!",
		int(calcSleeps(target)))
}
