package main

import (
	"encoding/json"
	"log"
	"os"
)

// a slice of users struct, output the country with the most users

// User represents a user record.
type User struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

const path = "users.json"

// getBiggestMarket takes in the slice of users and
// returns the biggest market.
func getBiggestMarket(users []User) (string, int) {
	countryMap := make(map[string]int)

	// create a map of countries and their counts
	for _, user := range users {
		countryMap[user.Country]++
	}

	// find the biggest market
	var biggestMarket string
	var biggestCount int

	for country, count := range countryMap {
		if count > biggestCount {
			biggestCount = count
			biggestMarket = country
		}
	}

	return biggestMarket, biggestCount
}

func main() {
	users := importData()
	country, count := getBiggestMarket(users)
	log.Printf("The biggest user market is %s with %d users.\n",
		country, count)
}

// importData reads the raffle entries from file and
// creates the entries slice.
func importData() []User {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data []User
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
