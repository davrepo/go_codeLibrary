package main

import (
	"log"
	"net"

	"example.com/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("failed to listen on port 8888: %v", err)
	}

	s := chat.Server{}

	gprcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(gprcServer, &s) // this is automatically generated method by protoc

	if err := gprcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server over port 8888: %v", err)
	}
}
