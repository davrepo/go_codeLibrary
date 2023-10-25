package chat

import (
	context "context"
	"log"
)

type Server struct {
	UnimplementedChatServiceServer // Embed the unimplemented server
}

func (s *Server) SendMessage(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &Message{Body: "Hello from the server!"}, nil
}
