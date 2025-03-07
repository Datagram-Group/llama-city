package chat

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	clients   = make(map[*websocket.Conn]bool) // list of WebSocket connections
	broadcast = make(chan []byte)              // channel for sending messages to clients
	mutex     sync.Mutex
)

// Handles broadcasting messages to all WebSocket clients
func HandleMessages() {
	for msg := range broadcast {
		mutex.Lock()
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Println("Error writing message to client:", err)
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}
