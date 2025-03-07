package websocket

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
)

var (
	clients = make(map[*websocket.Conn]bool) // list of WebSocket connections
	mutex   sync.Mutex
)

// WebSocket upgrader configuration
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Function to handle WebSocket connections
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade failed:", err)
		return
	}
	defer conn.Close()

	// Register new connection
	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	// Listen and process WebSocket data
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read failed:", err)
			break
		}
		fmt.Println("Received WebSocket message:", string(msg))

		// Send response in JSON format
		response := map[string]string{
			"status":  "ACK",
			"message": "Received message",
		}
		err = conn.WriteJSON(response)
		if err != nil {
			log.Println("Error writing message to client:", err)
			break
		}
	}

	// Remove connection when closed
	mutex.Lock()
	delete(clients, conn)
	mutex.Unlock()
}

func StartWebSocketServer() {
	// Get port from configuration
	wsPort := viper.GetString("websocket_port")
	http.HandleFunc("/ws", HandleWebSocket)

	fmt.Printf("llama-city WebSocket server is running on port %s...\n", wsPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", wsPort), nil); err != nil {
		log.Fatalf("Failed to start WebSocket server: %v", err)
	}
}
