// Go's map data structure
package main

import (
	"fmt"
)

func main() {
	stocks := map[string]float64{
		"AMZN": 2087.98,
		"GOOG": 2540.85,
		"MSFT": 287.70, // Must have trailing comma
	}

	fmt.Println(len(stocks)) // 3

	fmt.Println(stocks["MSFT"]) // 287.70

	// If key doesn't exist, get zero value
	fmt.Println(stocks["TSLA"]) // 0

	// Use two value to see if found
	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	} // TSLA not found

	// Insert new element
	stocks["TSLA"] = 822.12
	fmt.Println(stocks) // map[AMZN:2087.98 GOOG:2540.85 MSFT:287.7 TSLA:822.12]

	// Delete
	delete(stocks, "AMZN")
	fmt.Println(stocks) // map[GOOG:2540.85 MSFT:287.7 TSLA:822.12]

	// Single value "for" is on keys
	for key := range stocks {
		fmt.Println(key)
	} // GOOG MSFT TSLA

	// Double value "for" is key, value
	for key, value := range stocks {
		fmt.Printf("%s -> %.2f\n", key, value)
	} // GOOG -> 2540.85 MSFT -> 287.70 TSLA -> 822.12
}
