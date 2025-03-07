package grpc

import (
	"fmt"
	chat "llama-city/internal/proto"

	"log"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type server struct {
	chat.UnimplementedChatServiceServer
}

func StartGRPCServer() {
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &server{})

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
