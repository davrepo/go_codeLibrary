// JSON example
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// Request is a bank transactions
type Request struct {
	// field tags on the right side
	// tell the JSON decoder which struct field to use
	// for a given JSON key
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

// this is the JSON we're going to decode
var data = `
{
  "user": "Scrooge McDuck",
  "type": "deposit",
  "amount": 123.4
}
`

func main() {
	// ----------- DECODE JSON -----------
	// this simulates a io.Reader like a file/socket
	rdr := strings.NewReader(data)

	// Decode request
	dec := json.NewDecoder(rdr)

	var req Request // declare a previous defined struct
	if err := dec.Decode(&req); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}

	fmt.Printf("got: %+v\n", req)
	// got: {Login:Scrooge McDuck Type:deposit Amount:123.4}

	// ----------- ENCODE JSON -----------
	// Create response
	prevBalance := 1_000_000.0      // Loaded from database
	resp := map[string]interface{}{ // interface{} accepts any type
		"ok":      true,                     // boolean type
		"balance": prevBalance + req.Amount, // float64 type
	}

	// Encode respose
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(resp); err != nil {
		log.Fatalf("error: can't encode - %s", err)
	}
	// {"balance":1000123.4,"ok":true}
}
