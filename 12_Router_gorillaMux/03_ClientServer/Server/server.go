package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

/* RUN CODE BASE
(1) Run server.go in a terminal
(2) RUn client.go in another terminal
(2a) echo 'Hello, World!' | go run client.go set myKey
(2b) go run client.go get myKey
(2c) echo 'Value 2' | go run client.go set Key2
(2d) go run client.go list
(3) in Bash terminal, run curl requests:
(3a) curl -X POST -H "Content-Type: application/octet-stream" -d 'Value 3!' http://localhost:8080/client/Key3
(3b) curl -X GET http://localhost:8080/client
*/

const (
	maxSize = 10 * (1 << 20) // 10MB
)

var (
	db     = make(map[string][]byte) // memory-only "database"
	dbLock sync.RWMutex
)

func handleGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	// RLock() allows multiple readers to read the map
	// but only one writer can write to the map
	dbLock.RLock()
	defer dbLock.RUnlock()

	data, ok := db[key]
	if !ok {
		log.Printf("error get - unknown key: %q", key)
		http.Error(w, fmt.Sprintf("%q not found", key), http.StatusNotFound)
		return
	}

	// write the value to the response
	if _, err := w.Write(data); err != nil {
		log.Printf("error sending: %s", err)
	}
}

func handleSet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	defer r.Body.Close()
	rdr := io.LimitReader(r.Body, maxSize)
	// read everything from the LimitReader
	data, err := ioutil.ReadAll(rdr)
	if err != nil {
		log.Printf("read error: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbLock.Lock()
	defer dbLock.Unlock()
	// store the data in the map
	db[key] = data

	// return a JSON response of the key and the size of the data stored
	resp := map[string]interface{}{
		"key":  key,
		"size": len(data),
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("error sending: %s", err)
	}
}

func handleList(w http.ResponseWriter, r *http.Request) {
	// RLock() allows multiple readers to read the map
	dbLock.RLock()
	defer dbLock.RUnlock()

	keys := make([]string, 0, len(db))
	for key := range db {
		keys = append(keys, key)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(keys); err != nil {
		log.Printf("error sending: %s", err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/client/{key}", handleSet).Methods("POST")
	r.HandleFunc("/client/{key}", handleGet).Methods("GET")
	r.HandleFunc("/client", handleList).Methods("GET")
	http.Handle("/", r)

	addr := ":8080"
	log.Printf("server ready on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
