package constant

import (
	"sync"

	"github.com/gorilla/websocket"
)

var (
	Broadcast = make(chan []byte) // Channel for sending messages to Clients
	Mutex     sync.Mutex
	Clients   = make(map[*websocket.Conn]bool) // List of WebSocket connections
)

// Register WebSocket Client
func RegisterClient(conn *websocket.Conn) {
	Mutex.Lock()
	Clients[conn] = true
	Mutex.Unlock()
}

// Unregister WebSocket Client
func UnregisterClient(conn *websocket.Conn) {
	Mutex.Lock()
	delete(Clients, conn)
	Mutex.Unlock()
}
