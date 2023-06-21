package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// httpbin.org is a good website to test HTTP operations

func getRequestTest() {
	const httpbin = "https://httpbin.org/get"

	// Perform a GET operation
	resp, err := http.Get(httpbin)
	if err != nil {
		return
	}
	// The caller is responsible for closing the response
	defer resp.Body.Close()

	// We can access parts of the response to get information:
	fmt.Println("Status:", resp.Status)                // Status: 200 OK
	fmt.Println("Status Code:", resp.StatusCode)       // Status Code: 200
	fmt.Println("Protocol:", resp.Proto)               // Protocol: HTTP/2.0
	fmt.Println("Content length:", resp.ContentLength) // Content length: 272

	// Use a String Builder to build the content from bytes
	var sb strings.Builder

	// Read the content and write it to the builder
	content, _ := ioutil.ReadAll(resp.Body) // content is type []byte slice
	bytecount, _ := sb.Write(content)       // sb.Write returns the number of bytes written

	// Format the output
	fmt.Println(bytecount) // 272
	fmt.Println(sb.String())
	/*
		{
		  "args": {},
		  "headers": {
		    "Accept-Encoding": "gzip",
		    "Host": "httpbin.org",
		    "User-Agent": "Go-http-client/2.0",
		    "X-Amzn-Trace-Id": "Root=1-648077b4-1e82438b3b337bfa781d9866"
		  },
		  "origin": "213.237.66.37",
		  "url": "https://httpbin.org/get"
		}
	*/
}

func main() {
	// Execute a GET request
	getRequestTest()
}
