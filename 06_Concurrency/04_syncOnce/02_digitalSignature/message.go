package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
)

// execute some function exactly once and in a thread-safe way,
// which is useful when initializing shared resources that need to be created only once.
// shared resource is the 'sig' field of the Message struct,
// which represents the SHA-1 hash signature of the message's content.

type Message struct {
	Content string
	once    sync.Once
	sig     string // cached signature
}

// returns the digital signature of the message
func (m *Message) Signature() string {
	// Do() method checks if the function has been called before
	// If the function has already been called once, Do does nothing.
	m.once.Do(m.calcSig)
	return m.sig
}

func (m *Message) calcSig() {
	log.Printf("calculating signature")
	h := sha1.New()
	// take m.Content turns into a Reader, pass to io.Copy
	// which copies Reader (source) to h (destination)
	io.Copy(h, strings.NewReader(m.Content))
	// generate final hash value and convert to hex string
	// Sum returns SHA-1 checksum, nil means no extra data appended to hash input
	// %x formats has to lower-case hexadecimal characters
	m.sig = fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	m := Message{
		Content: "There is nothing more deceptive than an obvious fact.",
		// all other fields are initialized to their zero values
	}
	fmt.Println(m.Signature())
	// 2021/05/03 20:24:45 calculating signature
	// b931605bbbcdd058f9c33b11d7093fe8030b5413
	fmt.Println(m.Signature())
	// b931605bbbcdd058f9c33b11d7093fe8030b5413
}
