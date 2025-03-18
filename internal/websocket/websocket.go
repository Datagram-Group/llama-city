package websocket

import (
	"encoding/json"
	"fmt"
	"llama-city/internal/grpc"
	"llama-city/pkg/constant"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocket configuration
func createWebSocketUpgrader() websocket.Upgrader {
	return websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}

// Function to handle WebSocket connection
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := createWebSocketConnection(w, r)
	if err != nil {
		return
	}
	defer conn.Close()

	// Log successful WebSocket connection
	log.Println("WebSocket connection established with:", conn.RemoteAddr())

	// Register new connection
	constant.RegisterClient(conn)
	defer constant.UnregisterClient(conn)

	// Listen and process WebSocket data
	processWebSocketMessages(conn)
}

// Process WebSocket messages
func processWebSocketMessages(conn *websocket.Conn) {
	for {
		// Read message from WebSocket connection
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read failed:", err)
			break
		}

		// Log the received message
		fmt.Printf("Received message from client: %s\n", string(msg))

		// Define a struct for expected message structure
		var receivedMessage map[string]string
		err = json.Unmarshal(msg, &receivedMessage)
		if err != nil {
			log.Println("Error unmarshalling message:", err)
			break
		}

		// Check for "Ack message" content
		if ackMessage, exists := receivedMessage["Ack message"]; exists {
			if ackMessage == "Find" {
				go grpc.ResponseAckFind()
				// constant.WorkerFound <- true
			}
		}
	}
}

// Create WebSocket connection
func createWebSocketConnection(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader := createWebSocketUpgrader()
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade failed:", err)
		return nil, err
	}
	return conn, nil
}

// Function to start WebSocket server
func StartWebSocketServer(webSocketPort string) {
	http.HandleFunc("/ws", handleWebSocket)

	fmt.Printf("llama-city WebSocket server is running on port %s...\n", webSocketPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", webSocketPort), nil); err != nil {
		log.Fatalf("Failed to start WebSocket server: %v", err)
	}
}
