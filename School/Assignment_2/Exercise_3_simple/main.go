package main

import (
	"fmt"
	"sync"
	"time"
)

// Flags to represent TCP flags
type Flags struct {
	SYN bool
	ACK bool
	FIN bool
}

// Packet to represent a TCP packet
type Packet struct {
	Flags Flags
	Seq   int
	Ack   int
	Data  string
}

// Channels to simulate network communication through forwarder
var clientToForwarder = make(chan Packet, 10)
var forwarderToServer = make(chan Packet, 10)
var serverToForwarder = make(chan Packet, 10)
var forwarderToClient = make(chan Packet, 10)

var quit = make(chan bool)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go client(&wg)
	go server(&wg)
	go forwarder(&wg) // New forwarder goroutine

	wg.Wait()
}

func client(wg *sync.WaitGroup) {
	defer wg.Done()

	// TCP Handshake
	// Client sends SYN packet to Server
	clientToForwarder <- Packet{
		Flags: Flags{SYN: true},
		Seq:   1,
	}

	packetFromServer := <-forwarderToClient
	fmt.Printf("Client got: %s\n", packetString(packetFromServer))

	clientToForwarder <- Packet{
		Flags: Flags{ACK: true},
		Seq:   2,
		Ack:   packetFromServer.Seq + 1,
	}

	// Send 5 data messages
	for i := 1; i <= 5; i++ {
		msg := fmt.Sprintf("Data %d", i)
		clientToForwarder <- Packet{
			Seq:  i + 2,
			Data: msg,
		}
		packetFromServer = <-forwarderToClient
		fmt.Printf("Client got: %s\n", packetString(packetFromServer))
	}

	// Send FIN packet for termination
	clientToForwarder <- Packet{
		Flags: Flags{FIN: true},
		Seq:   8,
	}

	// Close the quit channel to signal forwarder to exit
	close(quit)
}

func server(wg *sync.WaitGroup) {
	defer wg.Done()

	packetFromClient := <-forwarderToServer
	fmt.Printf("Server got: %s\n", packetString(packetFromClient))

	// Send SYN-ACK to client
	serverToForwarder <- Packet{
		Flags: Flags{SYN: true, ACK: true},
		Seq:   1,
		Ack:   packetFromClient.Seq + 1,
	}

	packetFromClient = <-forwarderToServer
	fmt.Printf("Server got: %s\n", packetString(packetFromClient))

	// Receive 5 data messages
	var dataReceived []string
	for i := 1; i <= 5; i++ {
		packetFromClient = <-forwarderToServer
		fmt.Printf("Server got: %s\n", packetString(packetFromClient))
		dataReceived = append(dataReceived, packetFromClient.Data)

		// Acknowledge the data received
		serverToForwarder <- Packet{
			Flags: Flags{ACK: true},
			Seq:   i + 2,
			Ack:   packetFromClient.Seq + 1,
		}
	}

	packetFromClient = <-forwarderToServer // Receive FIN packet
	fmt.Printf("Server got: %s\n", packetString(packetFromClient))

	// Send ACK for FIN
	serverToForwarder <- Packet{
		Flags: Flags{ACK: true},
		Seq:   8,
		Ack:   packetFromClient.Seq + 1,
	}

	fmt.Println("Data received:", dataReceived)
}

func forwarder(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case packetFromClient := <-clientToForwarder:
			// Simulate delay
			time.Sleep(time.Millisecond * 10)
			forwarderToServer <- packetFromClient
		case packetFromServer := <-serverToForwarder:
			// Simulate delay
			time.Sleep(time.Millisecond * 10)
			forwarderToClient <- packetFromServer
		case <-quit: // When quit channel is closed
			return // Exit the forwarder goroutine
		}
	}
}

func packetString(p Packet) string {
	s := ""
	if p.Flags.SYN {
		s += "SYN "
	}
	if p.Flags.ACK {
		s += "ACK "
	}
	if p.Flags.FIN {
		s += "FIN "
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
