package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	XMLName    xml.Name `xml:"persondata"` // root level tag
	FirstName  string   `xml:"firstname"`
	LastName   string   `xml:"lastname"`
	Address    string   `xml:"addr"`
	Age        int      `xml:"age,attr"` // attr is attribute on root level tag
	FaveColors []string `xml:"favecolors"`
}

func main() {
	// Create some sample XML data
	xmldata := `
	<persondata age="29">
		<firstname>John</firstname>
		<lastname>Public</lastname>
        <addr>456 Everywhere Blvd</addr>
        <favecolors>Purple</favecolors>
        <favecolors>Yellow</favecolors>
        <favecolors>Green</favecolors>
	</persondata>
	`

	// Data will be decoded into a person struct
	var p person

	// Use the Unmarshal function to decode raw XML data
	xml.Unmarshal([]byte(xmldata), &p)
	fmt.Printf("%#v\n", p)
}
