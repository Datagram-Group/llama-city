package main

import (
	"llama-city/internal/config"
	"llama-city/internal/grpc"
	"llama-city/internal/message"
	"llama-city/internal/websocket"

	"log"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	go grpc.StartGRPCServer(config.GRPCServerPort)
	go websocket.StartWebSocketServer(config.WebSocketServerPort)
	go message.HandleMessages()

	select {}
}
