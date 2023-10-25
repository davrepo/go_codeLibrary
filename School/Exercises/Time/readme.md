# Overview
getTime function queries an NTP (Network Time Protocol) server to get the current time. This function performs a UDP-based NTP request to the specified host and receives a packet with time information. It calculates the elapsed time between sending the request and receiving the response to provide a more accurate time measurement.

# Parameters
host string: The hostname or IP address of the NTP server to connect to.
# Return Values
- *packet: A pointer to a packet structure containing the NTP response.
- int64: The current Unix time (in seconds) as received from the NTP server, adjusted for NTP epoch offset.
- error: An error object, which is nil if the operation was successful. Otherwise, it contains details about what went wrong.

# Dependencies
net: Go's standard library package for networking.
time: Go's standard library package for time manipulation.
binary: Go's standard library package for reading and writing binary data.
errors: Go's standard library package for handling errors.

# Example Usage
responsePacket, unixTime, err := getTime("time.google.com")
if err != nil {
    log.Fatal("Error getting time: ", err)
}
fmt.Printf("Received time: %v (Unix time: %d)\n", responsePacket, unixTime)
