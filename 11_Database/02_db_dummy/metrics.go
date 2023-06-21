package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

/*
To run the code base:
go run .
curl -X POST -H "Content-Type: application/json" -d @/c/Users/jackh/Dropbox/ITU\ DTU\ courses/Distributed\ Systems/go-essential-training-2446262/11_Database/02_db_dummy/metric.json http://localhost:8080/metric
*/

var (
	db *DB
)

// Metric is an application metric
type Metric struct {
	Time   time.Time `json:"time"`
	Host   string    `json:"host"`
	CPU    float64   `json:"cpu"`    // CPU load
	Memory float64   `json:"memory"` // MB
}

func main() {
	http.HandleFunc("/metric", handleMetric)

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}

	log.Printf("server ready on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func handleMetric(w http.ResponseWriter, r *http.Request) {
	// if request is not POST, then error
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	var m Metric
	// limit the amount of data you read from the request
	const maxSize = 1 << 20 // MB
	dec := json.NewDecoder(io.LimitReader(r.Body, maxSize))
	// decode the request body into a Metric struct m
	if err := dec.Decode(&m); err != nil {
		log.Printf("error decoding: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := db.Add(m)
	log.Printf("metric: %+v (id=%s)", m, id)

	// send response to client, reply with the ID of the metric
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]interface{}{
		"id": id,
	}
	// encode resp into JSON and write it to the response w
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		// if there is an error, we cant change HTTP status code
		// because we already wrote the header, so we just log the error
		log.Printf("error reply: %s", err)
	}
}
