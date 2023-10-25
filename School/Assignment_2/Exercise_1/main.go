package main

import (
	"fmt"
	"sync"
)

// Flags to represent TCP flags
type Flags struct {
	SYN bool
	ACK bool
}

// Packet to represent a TCP packet
type Packet struct {
	Flags Flags
	Seq   int
	Ack   int
	Data  string
}

// Channels to simulate network communication
var clientToServer = make(chan Packet, 1)
var serverToClient = make(chan Packet, 1)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go client(&wg)
	go server(&wg)

	wg.Wait()
}

func client(wg *sync.WaitGroup) {
	defer wg.Done()

	// Client sends SYN packet to Server
	clientToServer <- Packet{
		Flags: Flags{SYN: true},
		Seq:   1,
	}

	// Client waits for SYN-ACK from server
	packetFromServer := <-serverToClient
	fmt.Printf("Client got: %s\n", packetString(packetFromServer))

	// Client sends ACK to server
	clientToServer <- Packet{
		Flags: Flags{ACK: true},
		Seq:   2,
		Ack:   packetFromServer.Seq + 1,
	}
}

func server(wg *sync.WaitGroup) {
	defer wg.Done()

	// Server waits for SYN from client
	packetFromClient := <-clientToServer
	fmt.Printf("Server got: %s\n", packetString(packetFromClient))

	// Server sends SYN-ACK to client
	serverToClient <- Packet{
		Flags: Flags{SYN: true, ACK: true},
		Seq:   1,
		Ack:   packetFromClient.Seq + 1,
	}

	// Server waits for ACK from client
	packetFromClient = <-clientToServer
	fmt.Printf("Server got: %s\n", packetString(packetFromClient))

	// Server sends data to client
	clientToServer <- Packet{
		Seq:  packetFromClient.Ack + 1,
		Data: "Data Transmitted!",
	}

	fmt.Printf("Server got: seq=%d data=%s\n", packetFromClient.Seq+1, "Data Transmitted!")
}

func packetString(p Packet) string {
	s := ""
	if p.Flags.SYN {
		s += "syn "
	}
	if p.Flags.ACK {
		s += "ack "
	}
	s += fmt.Sprintf("seq=%d ", p.Seq)
	if p.Flags.ACK {
		s += fmt.Sprintf("ack=%d ", p.Ack)
	}
	if p.Data != "" {
		s += fmt.Sprintf("data=%s", p.Data)
	}
	return s
}
