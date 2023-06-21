package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
1.	defines a struct type Quantity that has two fields.
2.	It then defines a method MarshalJSON on Quantity that satisfies the json.Marshaler interface.
This interface can be implemented by types that need to control their own serialization in JSON.
When you call json.Marshal on a value of a type that implements this interface,
it will use this method instead of the default behavior.
3.	In the MarshalJSON method, it uses fmt.Sprintf to create a string that concatenates the Value and Unit fields.
Sprintf formats a string according to a format specifier (in this case, %f%s means float and string),
using the provided values. It does not print this string anywhere; it only returns it.
4.	It then uses json.Marshal to convert this string to a JSON value, which is a byte slice ([]byte).
This is the standard format for JSON data in Go.
5.	In main, it creates a Quantity value, and encodes it to JSON using json.NewEncoder(os.Stdout).Encode(&q).
This writes the JSON output to standard output (os.Stdout).
The resulting JSON is a single JSON string: "1.780000meter".
The method Encode calls the MarshalJSON method that was defined for the Quantity type.
The MarshalJSON method here is used to customize the way Quantity values are encoded to JSON.
Instead of the default behavior, which would produce a JSON object with Value and Unit fields
(like {"Value":1.78,"Unit":"meter"}), it produces a single JSON string with the value and unit concatenated.
*/

// Quantity is combination of value and unit (e.g. 2.7cm)
type Quantity struct {
	Value float64
	Unit  string
}

// MarshalJSON implements the json.Marshaler interface
// instead of encoding JSON object w/ 2 fields, we encode a single string
// Example encoding: "42.195km"
func (q *Quantity) MarshalJSON() ([]byte, error) {
	if q.Unit == "" {
		return nil, fmt.Errorf("empty unit")
	}
	// Sprintf returns a string
	text := fmt.Sprintf("%f%s", q.Value, q.Unit)
	// json.Marshal converts string to []byte slice
	return json.Marshal(text)
}

func main() {
	q := Quantity{1.78, "meter"}
	json.NewEncoder(os.Stdout).Encode(&q) // "1.780000meter"
}
