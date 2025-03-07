package chat

import (
	"context"
	"encoding/json"
	"fmt"
	chat "llama-city/internal/proto"

	"log"
)

type Message struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

// server is the structure for gRPC server
type Server struct {
	chat.UnimplementedChatServiceServer
}

// Handles the SendMessage gRPC request
func (s *Server) SendMessage(ctx context.Context, req *chat.ChatRequest) (*chat.ChatResponse, error) {
	// Handles the message and sends it via WebSocket
	messageJson, err := handleGRPCMessage(req)
	if err != nil {
		return nil, err
	}
	broadcast <- messageJson

	return &chat.ChatResponse{AckMessage: "ACK Find"}, nil
}

func handleGRPCMessage(req *chat.ChatRequest) ([]byte, error) {
	var message Message
	message.Model = req.Model
	for _, msg := range req.Messages {
		message.Messages = append(message.Messages, struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{Role: msg.Role, Content: msg.Content})
	}
	messageJson, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshaling message:", err)
		return nil, err
	}
	fmt.Printf("messageJson: %s\n", messageJson)
	return messageJson, nil
}
