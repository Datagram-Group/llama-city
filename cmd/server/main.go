package main

import (
	"llama-city/internal/chat"
	"llama-city/internal/config"
	"llama-city/pkg/grpc"
	"llama-city/pkg/websocket"
)

func main() {

	config.LoadConfig()

	go grpc.StartGRPCServer()
	go websocket.StartWebSocketServer()
	go chat.HandleMessages()

	select {}
}
