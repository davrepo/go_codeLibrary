package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

// 10 shares of MSFT for $99.98
var transRe = regexp.MustCompile(`(\d+) shares of ([A-Z]+) for \$(\d+(\.\d+)?)`)

type Transaction struct {
	Symbol string
	Volume int
	Price  float64
}

func parseLine(line string) (Transaction, error) {
	matches := transRe.FindStringSubmatch(line) // return []string
	if matches == nil {
		return Transaction{}, fmt.Errorf("bad line: %q", line)
	}
	var t Transaction
	t.Symbol = matches[2]
	t.Volume, _ = strconv.Atoi(matches[1])
	t.Price, _ = strconv.ParseFloat(matches[3], 64)
	return t, nil
}

func main() {
	line := "10 shares of MSFT for $99.98"
	t, err := parseLine(line)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", t) // {Symbol:MSFT Volume:10 Price:99.98}
}
