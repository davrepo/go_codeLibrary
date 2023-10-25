a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?

Packages: fmt, sync. 

type Packet struct {
	Flags Flags
	Seq   int
	Ack   int
	Data  string
}

The primary data structure used for transmitting data and meta-data is the Packet struct
Flags: This is another struct that represents the TCP flags. In this simulation, two flags are defined: SYN and ACK.
Seq: Represents the sequence number of the packet.
Ack: Represents the acknowledgment number of the packet.
Data: Represents the data content of the packet (used in the server function when sending a "Data Transmitted!" message).

The Packet struct encapsulates both data (the Data field) and meta-data (the Flags, Seq, and Ack fields) being transmitted between the client and server via the forwarders.


b) Does your implementation use threads or processes? Why is it not realistic to use threads?

The implementation uses goroutines. It is neither a thread nor process. Multiple goroutines can run on a single OS-level thread. Goroutines are multiplexed onto a smaller number of OS-level threads by the Go scheduler.

Thread on the other hand is the smallest unit of CPU execution. Each thread has its own registers, stack, and local storage. A process can have multiple threads, and all these threads share the process's memory and resources. The primary difference between threads and processes is the isolation level: processes are more isolated from each other, while threads within the same process can directly communicate and share memory. So the implemention is with Goroutines, which is neither a thread nor process.

For the second part of the question. (1) Creating an OS-level thread for each simulated handshake would be inefficient, especially if you have a large number of handshakes. Threads have more overhead in terms of memory and context-switching compared to goroutines. (2) Operating systems have limits on the number of threads that can be created, and you'd hit this ceiling much faster than with goroutines. 

Using a separate thread for each simulated connection would not scale well and isn't necessary for the purpose of this simulation.


c) In case the network changes the order in which messages are delivered, how would you handle message re-ordering?

The order of received packets is ensured by the Sequence (Seq) and Acknowledgment (Ack) numbers. If the network reorders the packets, the receiver can use these numbers to reorder them correctly.
(1) Maintain a buffer at the receiver side. As out-of-order packets arrive, place them in the buffer according to their sequence numbers.
(2) Send acknowledgments for the highest sequence number of a packet that has been received in order.
(3) Once the missing packets arrive and all packets up to a certain sequence number are available, deliver them to the application in the correct order.
Implement a timeout mechanism. If the expected packet doesn't arrive within a certain time frame, the receiver can request a retransmission.
Sender also have its own timeout mechanism. If it doesn't receive an acknowledgment for a sent packet within a certain time frame, it should consider the packet as lost and retransmit it.
Sliding Window Protocol: receiver has a "window" of acceptable sequence numbers for incoming packets. This window slides forward as packets are received and acknowledged.


d) In case messages can be delayed or lost, how does your implementation handle message loss?

message loss is simulated by introducing a 20% chance for the forwarder function to drop a packet. If a packet is dropped, both the client and server are designed to detect this loss.

Client-side:
The client sends a packet with the SYN flag set.
It then waits for a packet from the server with both the SYN and ACK flags set.
If the client doesn't receive this packet (because it got lost in transit), it will be stuck waiting. The same applies to the final ACK packet the client sends; if it doesn't get a response, it knows something went wrong.

Server-side:
The server waits for the initial packet from the client with the SYN flag set.
If it doesn't receive this packet, it will be stuck waiting.
After sending its SYN-ACK response, the server waits for an ACK packet from the client. If it doesn't receive this packet, it again knows that something went wrong.

In both cases, the client and server will eventually time out and exit the program.

e) Why is the 3-way handshake important?
It is important for establishing mutual agreement. (1) It ensures that both client and server have agreed to establish a connection. (2) The handshake allows both parties to exchange and acknowledge initial sequence numbers. This guarantees that subsequent data transfers are orderly and that both sides are synchronized in their understanding of the sequence of the data. (3) By requiring each party to acknowledge receipt of initial connection parameters (like sequence numbers), the handshake ensures that the communication channel is reliable. It verifies that both sides can send and receive data correctly. (4) helps protect against scenarios where delayed packets from an old connection might be mistaken for a new connection request.

Realiability: it guarantees the delivery of data in the correct order without duplicates. 