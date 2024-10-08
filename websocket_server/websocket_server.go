package websocket_server

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // Connected clients
var broadcast = make(chan ProductUpdate)     // Broadcast channel for product updates
var mu sync.Mutex

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections by default
		return true
	},
}

// ProductUpdate struct to represent updates
type ProductUpdate struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

// WebSocket endpoint
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// Register the new client
	mu.Lock()
	clients[ws] = true
	mu.Unlock()

	// Wait for new product updates and send to clients
	for {
		var product ProductUpdate
		err := ws.ReadJSON(&product)
		if err != nil {
			log.Printf("Error: %v", err)
			mu.Lock()
			delete(clients, ws)
			mu.Unlock()
			break
		}

		// Send the product update to the broadcast channel
		broadcast <- product
	}
}

// Handle broadcasting product updates to all clients
func HandleBroadcasts() {
	for {
		// Grab the product from the broadcast channel
		product := <-broadcast

		// Send the product update to every client
		mu.Lock()
		for client := range clients {
			err := client.WriteJSON(product)
			if err != nil {
				log.Printf("Error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
		mu.Unlock()
	}
}
