package message

import (
	"encoding/json"
	"llama-city/internal/models"
	chat "llama-city/internal/proto"
	"llama-city/pkg/constant"

	"log"
)

// Process messages from gRPC
func ProcessGRPCMessage(req *chat.ChatRequest) ([]byte, error) {
	var message models.Message
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

	return messageJson, nil
}

// Send message to all WebSocket Clients
func sendToAllClients(msg []byte) {
	constant.Mutex.Lock()
	for client := range constant.Clients {
		err := client.WriteJSON(msg)
		if err != nil {
			log.Println("Error writing message to client:", err)
			client.Close()
			delete(constant.Clients, client)
		}
	}
	constant.Mutex.Unlock()
}

// Function to broadcast message to all WebSocket Clients
func HandleMessages() {
	for msg := range constant.Broadcast {
		sendToAllClients(msg)
	}
}
