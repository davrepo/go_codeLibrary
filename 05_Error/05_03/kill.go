package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

// read pid from a file
// convert it from a string to an integer
// then print out "killing server with pid"
func killServer(pidFile string) error {
	// Error: can't open pid file
	file, err := os.Open(pidFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Error: can't read pid
	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return errors.Wrap(err, "can't read pid")
	}

	// Simulate kill
	fmt.Printf("killing server with pid=%d\n", pid)

	if err := os.Remove(pidFile); err != nil {
		log.Printf("warning: can't remove pid file: %s", err)
	}

	return nil // return nil signals that there was no error
}

func main() {
	if err := killServer("server.pid"); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
