package main

import (
	"context"
	"log"

	"example.com/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	creds := insecure.NewCredentials()
	conn, err := grpc.Dial(":8888", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	} 
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from the client!",
	}

	response, err := c.SendMessage(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SendMessage: %s", err)
	}

	log.Printf("Response from server: %s", response.Body)
}
