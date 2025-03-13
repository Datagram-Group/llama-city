package websocket

import (
	"fmt"
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

	// Register new connection
	constant.RegisterClient(conn)
	defer constant.UnregisterClient(conn)

	// Listen and process WebSocket data
	processWebSocketMessages(conn)
}

// Process WebSocket messages
func processWebSocketMessages(conn *websocket.Conn) {
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read failed:", err)
			break
		}
		// fmt.Println("Received WebSocket message:", string(msg))

		// Send response in JSON format
		err = sendWebSocketResponse(conn)
		if err != nil {
			break
		}
	}
}

// Send WebSocket response
func sendWebSocketResponse(conn *websocket.Conn) error {
	response := map[string]string{
		"status":  "ACK",
		"message": "Received message",
	}
	err := conn.WriteJSON(response)
	if err != nil {
		log.Println("Error writing message to client:", err)
		return err
	}
	return nil
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
