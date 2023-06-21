package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Metric is an application metric
type Metric struct {
	Time   time.Time `json:"time"`
	CPU    float64   `json:"cpu"`    // CPU load
	Memory float64   `json:"memory"` // MB
}

func main() {
	m := Metric{
		Time:   time.Now(),
		CPU:    0.23,
		Memory: 87.32,
	}
	if err := postMetric(m); err != nil {
		log.Fatal(err)
	}
	// 2021/04/30 17:53:15 GOT: {Time:2021-04-30 17:53:14.437272 +0300 IDT CPU:0.23 Memory:87.32}
}

// Marshal Metric m struct => JSON
// then send JSON to httpbin.org/post
// get a response from server and decode it back to Metric struct
func postMetric(m Metric) error {
	// marshal Metric struct to JSON []byte slice
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	// if request doesnt complee within 3 seconds, cancel it
	// context.WithTimeout returns a copy of parent context with a new Done channel and a new cancel function
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// cancel function cancels the context, releases resources associated with it
	// this is for cleaning up incase request completes before the timeout 3 seconds is reached
	defer cancel()

	const url = "https://httpbin.org/post"
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(data))
	if err != nil {
		return err
	}
	// this tells the server that the body of the request will be in JSON format
	req.Header.Set("Content-Type", "application/json")

	// send HTTP request and returns a HTTP response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %d %s", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()

	// limit max size of response to 1MB to prevent consuming too much memory
	// if server returns a large response
	const maxSize = 1 << 20 // 1MB
	r := io.LimitReader(resp.Body, maxSize)
	var reply struct {
		JSON Metric // JSON is a field of type Metric
	}
	// decodes server response from JSON to Metric struct
	if err := json.NewDecoder(r).Decode(&reply); err != nil {
		return err
	}
	log.Printf("GOT: %+v\n", reply.JSON)
	return nil
}
