package main

import (
	"flag"
	"llama-city/internal/config"
	"llama-city/internal/grpc"
	"llama-city/internal/message"
	"llama-city/internal/websocket"

	"log"
)

func main() {
	configPath := flag.String("config", "./config.yaml", "Path to the config file")

	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	go grpc.StartGRPCServer(cfg.GRPCServerPort)
	go websocket.StartWebSocketServer(cfg.WebSocketServerPort)
	go message.HandleMessages()

	select {}
}
