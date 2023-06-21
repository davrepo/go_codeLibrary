package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const apiBase = "http://localhost:8080/client"

func get(key string) error {
	// url is apiBase/client/key
	url := fmt.Sprintf("%s/%s", apiBase, key)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %d %s", resp.StatusCode, resp.Status)
	}

	// copy the response body to os.Stdout
	_, err = io.Copy(os.Stdout, resp.Body)
	return err
}

func list() error {
	// list() does not need a key b/c list all
	resp, err := http.Get(apiBase)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %d %s", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()

	var keys []string
	if json.NewDecoder(resp.Body).Decode(&keys); err != nil {
		return err
	}
	for _, key := range keys {
		fmt.Println(key)
	}
	return nil
}

func set(key string) error {
	// set() construct the URL with apiBase/client/key
	// and then POST the data from os.Stdin to that URL.
	url := fmt.Sprintf("%s/%s", apiBase, key)
	// create http.Post() to url, with content-type "application/octet-stream"
	// which means that the body of the request is not a JSON document
	// but a stream of bytes (binary data).
	// the body of the request is comimg from os.Stdin
	resp, err := http.Post(url, "application/octet-stream", os.Stdin)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %d %s", resp.StatusCode, resp.Status)
	}

	var reply struct {
		Key  string
		Size int
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&reply); err != nil {
		return err
	}

	// print the key and the size of the data stored to stdout
	// set() does not store data to server by itself,
	// storage is via server's handleSet() function through http.Post()
	fmt.Printf("%s: %d bytes\n", reply.Key, reply.Size)
	return nil
}

func main() {
	// for displaying help message, with -h or --help
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: kv get|set|list [key]")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatalf("error: wrong number of arguments")
	}

	switch flag.Arg(0) {
	case "get":
		key := flag.Arg(1)
		if key == "" {
			log.Fatalf("error: missing key")
		}
		if err := get(key); err != nil {
			log.Fatal(err)
		}
	case "set":
		key := flag.Arg(1)
		if key == "" {
			log.Fatalf("error: missing key")
		}
		if err := set(key); err != nil {
			log.Fatal(err)
		}
	case "list":
		if err := list(); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("error: unknown command: %q", flag.Arg(0))
	}
}
