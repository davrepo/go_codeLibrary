// struct demo
package main

import (
	"fmt"
	"time"
)

// Budget is a budgest for campaign
type Budget struct {
	CampaignID string
	Balance    float64 // USD
	Expires    time.Time
}

// TimeLeft() is a method of Budget
func (b Budget) TimeLeft() time.Duration {
	// returns Expiration time - current time (subtract)
	return b.Expires.Sub(time.Now().UTC())
}

// Update() is a method of Budget
// *Budget is a pointer to Budget - so can modify the value without copying
func (b *Budget) Update(sum float64) {
	b.Balance += sum
	// b is a pointer to Budget
}

func main() {
	// create a new budget, expires in 7 days
	b := Budget{"Kittens", 22.3, time.Now().Add(7 * 24 * time.Hour)}
	fmt.Println(b.TimeLeft())
	// 168h0m0s

	b.Update(10.5) // update the budget balance with pointer
	fmt.Println(b.Balance)
	// 32.8				// 22.3 + 10.5

}
