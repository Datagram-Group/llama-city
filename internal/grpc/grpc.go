package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"llama-city/internal/message"
	chat "llama-city/internal/proto"
	"llama-city/pkg/constant"

	"google.golang.org/grpc"
)

type server struct {
	chat.UnimplementedChatServiceServer
}

// Function to handle gRPC request
func (s *server) SendMessage(ctx context.Context, req *chat.ChatRequest) (*chat.ChatResponse, error) {
	message, err := message.ProcessGRPCMessage(req)
	if err != nil {
		return nil, err
	}

	// Add message to channel to broadcast to all WebSocket clients
	constant.Broadcast <- message

	return &chat.ChatResponse{AckMessage: "ACK Find Worker"}, nil
}

func ResponseAckFind() *chat.ChatResponse {
	return &chat.ChatResponse{AckMessage: "ACK Find Worker"}
}

// Function to start gRPC server
func StartGRPCServer(gRPCServerPort string) {
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &server{})

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCServerPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Printf("llama-city gRPC server is running on port %s...\n", gRPCServerPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
