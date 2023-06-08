package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name       string   `json:"fullname"`
	Address    string   `json:"addr"`
	Age        int      `json:"age"`
	FaveColors []string `json:"favecolors"`
}

func decodeExample() {
	// declare some sample JSON data to decode
	data := []byte(`
		{
			"fullname" : "John Q Public",
			"addr" : "987 Main St",
			"age": 45,
			"favecolors" : ["Purple","White","Gold"]
		}
	`)

	// JSON will be decoded into a person struct
	var p person

	// test to see if the JSON is valid
	valid := json.Valid(data)
	if valid {
		// Unmarhsal - decodes JSON into data structure
		json.Unmarshal(data, &p) // decode data into p
		fmt.Printf("%#v\n", p)
	}

	// if JSON structure is unknown, decoded into a map
	var m map[string]interface{}

	// Unmarshal into a map
	json.Unmarshal(data, &m)
	fmt.Printf("%#v\n", m)
	for k, v := range m {
		fmt.Printf("key (%v), value (%T : %v)\n", k, v, v)
	}
}

func main() {
	// Decode JSON into Go structs
	decodeExample()
}
