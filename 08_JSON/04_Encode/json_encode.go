package main

import (
	"encoding/json"
	"fmt"
)

// convert data structure to JSON
type person struct {
	// json tag is used to specify the name of the field in the JSON string
	Name    string `json:"fullname"`
	Address string `json:"addr"`
	// if you do not want Age to appear in JSON, tag it with "-", so `json:"-"`
	Age int `json:"age"`
	// omit empty values means if the value is empty, it will not be included in the JSON string
	FaveColors []string `json:"favecolors,omitempty"`
}

func encodeExample() {
	// create some people data
	people := []person{
		{"Jane Doe", "123 Anywhere Street", 35, nil},
		{"John Public", "456 Everywhere Blvd", 29, []string{"Purple", "Yellow", "Green"}},
	}

	// Marshal - convert data structure to JSON
	result, err := json.Marshal(people)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", result)
	// [{"fullname":"Jane Doe","addr":"123 Anywhere Street","age":35},{"fullname":"John Public","addr":"456 Everywhere Blvd","age":29,"favecolors":["Purple","Yellow","Green"]}]

	// MarshalIndent - format the JSON string with indentation (easier for human to read)
	result1, err := json.MarshalIndent(people, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", result1)
}

func main() {
	// Encode Go data as JSON
	encodeExample()
}
