package main

import (
	"encoding/json"
	"fmt"
)

// dealing with zero / wrong values in JSON
// JSON => Go struct

type LineItem struct {
	SKU      string
	Price    float64
	Discount float64
	Quantity int
}

// NewLineItem initializes a LineItem struct w/ Quantity = 1
// and all other fields set to their zero values
func NewLineItem() LineItem {
	return LineItem{
		Quantity: 1,
	}
}

func unmarshalLineItem(data []byte) (LineItem, error) {
	li := NewLineItem()
	// Unmarshal data JSON into li struct
	if err := json.Unmarshal(data, &li); err != nil {
		return LineItem{}, nil
	}

	if li.Quantity < 1 {
		return LineItem{}, fmt.Errorf("bad quantity")
	}

	return li, nil
}

func main() {
	// JSON data missing Quantity field
	data := []byte(`{"sku": "x3xs", "price": 1.2}`)
	li, err := unmarshalLineItem(data)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Printf("%#v\n", li)
	}
	// main.LineItem{SKU:"x3xs", Price:1.2, Discount:0, Quantity:1}

	// JSON data has wrong value for Quantity field
	data = []byte(`{"sku": "x3xs", "price": 1.2, "quantity": 0}`)
	li, err = unmarshalLineItem(data)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Printf("%#v\n", li)
	}
	// ERROR: bad quantity
}
