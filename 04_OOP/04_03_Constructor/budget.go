// struct demo
package main

import (
	"fmt"
	"time"
)

type Budget struct {
	CampaignID string
	Balance    float64 // USD
	Expires    time.Time
}

// NewBudget is constructor function
// returns a pointer to Budget and an error
func NewBudget(campaignID string, balance float64, expires time.Time) (*Budget, error) {
	// data validation
	if campaignID == "" {
		return nil, fmt.Errorf("empty campaignID")
	}

	if balance <= 0 {
		return nil, fmt.Errorf("balance must be bigger than 0")
	}

	if expires.Before(time.Now()) {
		return nil, fmt.Errorf("bad expiration date")
	}

	b := Budget{
		CampaignID: campaignID,
		Balance:    balance,
		Expires:    expires,
	}
	// New function returns a pointer to Budget and an error
	// &b creates a pointer to the newly created Budget struct.
	return &b, nil
}

func main() {
	expires := time.Now().Add(7 * 24 * time.Hour)

	b1, err := NewBudget("puppies", 32.2, expires)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Printf("%#v\n", b1) // %#v - print the struct
	}
	// &main.Budget{CampaignID:"puppies", Balance:32.2, Expires:time.Date(2023, time.June, 4, 16, 27, 1, 58843200, time.Local)}

	b2, err := NewBudget("kittens", -3.2, expires)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Printf("%#v\n", b2)
	}
	// ERROR: balance must be bigger than 0
}
