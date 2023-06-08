// struct demo
package main

import (
	"fmt"
	"time"
)

// Upper case is public
// lower case is private
type Budget struct {
	Id      string
	Balance float64 // USD
	Expires time.Time
}

func main() {
	// := is shorthand for || var b2 Budget = Budget{...}
	b1 := Budget{"Kittens", 22.3, time.Now().Add(7 * 24 * time.Hour)}
	fmt.Println(b1)
	// {Kittens 22.3 2023-06-11 19:52:45.612956 +0200 CEST m=+604800.004269201}

	fmt.Printf("%#v\n", b1) // print with field names
	// main.Budget{Id:"Kittens", Balance:22.3, Expires:time.Date(2023, time.June, 11, 19, 52, 45, 612956000, time.Local)}

	fmt.Println(b1.Id) // access a field
	// Kittens

	// pass value by attribute name // order doesn't matter
	// Expires type is the zero value for time.Time
	b2 := Budget{
		Balance: 19.3,
		Id:      "puppies",
	}
	fmt.Printf("%#v\n", b2)
	// main.Budget{Id:"puppies", Balance:19.3, Expires:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}

	// define all fileds with zero values
	var b3 Budget
	fmt.Printf("%#v\n", b3)
	// main.Budget{Id:"", Balance:0, Expires:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}
}
