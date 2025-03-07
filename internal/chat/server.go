package chat

import (
	"fmt"
	chat "llama-city/internal/proto"
	"log"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func StartGRPCServer() {
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &Server{})

	grpcPort := viper.GetString("grpc_port")
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Printf("gRPC server running on port %s...\n", grpcPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
