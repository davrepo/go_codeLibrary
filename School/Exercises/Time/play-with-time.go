package main

//kudos to https://github.com/beevik/ntp/blob/master/ntp.go for inspiration

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"time"
)

const (
	defaultNtpVersion = 4                //  default NTP version to use
	defaultTimeout    = 15 * time.Second // timeout for the UDP connection, set to 15 seconds.
)

var (
	// difference between the Unix and NTP epoch
	ntpEpochOffset = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC).Unix() - time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
)

// structure of an NTP packet.
type packet struct {
	Settings       uint8  // leap yr indicator, ver number, and mode
	Stratum        uint8  // stratum of local clock
	Poll           int8   // poll exponent
	Precision      int8   // precision exponent
	RootDelay      uint32 // root delay
	RootDispersion uint32 // root dispersion
	ReferenceID    uint32 // reference id
	RefTimeSec     uint32 // reference timestamp sec
	RefTimeFrac    uint32 // reference timestamp fractional
	OrigTimeSec    uint32 // origin time secs
	OrigTimeFrac   uint32 // origin time fractional
	RxTimeSec      uint32 // receive time secs
	RxTimeFrac     uint32 // receive time frac
	TxTimeSec      uint32 // transmit time secs
	TxTimeFrac     uint32 // transmit time frac
}

func main() {
	// start and t capture the current time using different clocks.
	start := time.Now() // Wall clock reading
	t := time.Now()

	// Monotonic clock reading
	elapsed := t.Sub(start) // Calculates time elapsed between start and t. Sub() is a method of the time.Time type. Monotonic clocks are useful for measuring time intervals because they are unaffected by changes to the system clock. Wall clocks are affected by changes to the system clock, such as when the user changes the time zone or the system clock is adjusted to synchronize with a time server.
	fmt.Printf("Start time %d elapsed %d\n", start, elapsed)
	fmt.Printf("Local %s\n Standard (ISO 8601) time %s\n UTC time %s\n Unix time %d\n", t.Local().String(), t, t.UTC(), t.Unix())

	fmt.Println("Getting time from NTP server ... ")
	msg, ntp, err := getTime("dk.pool.ntp.org") // get NTP time and stores the returned values in msg, ntp, and err.
	// msg is a pointer to a packet struct, ntp is the time in seconds since the Unix epoch, and err is an error value.
	if err != nil {
		panic(err)
	}

	fmt.Printf("Msg %+v\nNTP time %d as Unix time \n.....", msg, ntp)
}

func getTime(host string) (*packet, int64, error) {
	/*
		# Overview
		getTime function queries an NTP (Network Time Protocol) server to get the current time. This function performs a UDP-based NTP request to the specified host and receives a packet with time information. It calculates the elapsed time between sending the request and receiving the response to provide a more accurate time measurement.

		# Parameters
		host string: The hostname or IP address of the NTP server to connect to.
		# Return Values
		- *packet: A pointer to a packet structure containing the NTP response.
		- int64: The current Unix time (in seconds) as received from the NTP server, adjusted for NTP epoch offset.
		- error: An error object, which is nil if the operation was successful. Otherwise, it contains details about what went wrong.

		# Example Usage
		responsePacket, unixTime, err := getTime("time.google.com")
		if err != nil {
				log.Fatal("Error getting time: ", err)
		}
		fmt.Printf("Received time: %v (Unix time: %d)\n", responsePacket, unixTime)
	*/

	radr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(host, "123")) // Resolves UDP address of NTP server.
	var ladr *net.UDPAddr                                                 // Defines a variable for the local address
	if err != nil {
		return nil, 0, err
	}
	conn, err := net.DialUDP("udp", ladr, radr) // Dials the NTP server using UDP.
	if err != nil {
		return nil, 0, err
	}
	defer conn.Close() // Ensures the connection is closed when the function exits.

	conn.SetDeadline(time.Now().Add(defaultTimeout)) // Sets a deadline for the UDP connection.

	req := &packet{Settings: 0x1B} // Creates a new NTP request packet with specific settings.
	xmitTime := time.Now()         // Records the current time before sending the NTP request.

	// query NTP
	err = binary.Write(conn, binary.BigEndian, req) // Writes the NTP request packet to the UDP connection.
	if err != nil {
		return nil, 0, err
	}
	// Receive the response.
	rsp := new(packet)                             // Creates a new packet to store the NTP response
	err = binary.Read(conn, binary.BigEndian, rsp) // Reads the NTP response from the UDP connection into rsp.
	if err != nil {
		return nil, 0, err
	}
	delta := time.Since(xmitTime) // Calculates the time elapsed since xmitTime was recorded.
	if delta < 0 {                // Checks if the elapsed time is negative, which shouldn't happen with a monotonic clock.
		// The local system may have had its clock adjusted since it
		// sent the query. In go 1.9 and later, time.Since ensures
		// that a monotonic clock is used, so delta can never be less
		// than zero. In versions before 1.9, a monotonic clock is
		// not used, so we have to check.
		return nil, 0, errors.New("client clock ticked backwards")
	}

	secs := int64(rsp.TxTimeSec) - ntpEpochOffset // Converts the received NTP time to Unix time.

	return rsp, secs, nil

}
