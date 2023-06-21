package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func postRequestTest() {
	const httpbin = "https://httpbin.org/post"

	// POST function
	// func Post(url, contentType string, body io.Reader) (resp *Response, err error)

	// POST operation using Post (Method 1)
	reqBody := strings.NewReader(`
	{
		"field1" : "This is field 1"
		"field2" : 250
	}
	`)
	resp, err := http.Post(httpbin, "application/json", reqBody)
	if err != nil {
		return
	}
	content, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Printf("%s\n", content)

	// POST operation using PostForm (Method 2)
	// let you send form encoded key/value pairs instead of raw JSON
	data := url.Values{} // Values is a map[string][]string
	data.Add("field1", "field added via Values")
	data.Add("field2", "300")
	resp1, err := http.PostForm(httpbin, data)
	if err != nil {
		return
	}

	content1, _ := ioutil.ReadAll(resp1.Body)
	defer resp.Body.Close()
	fmt.Printf("%s\n", content1)
}

func main() {
	// Execute a POST
	postRequestTest()
}
