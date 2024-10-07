package websocket_server

// import (
// 	"log"
// 	"net/http"
// 	"sync"

// 	"github.com/gorilla/websocket"
// )

// var clients = make(map[*websocket.Conn]bool) // Connected clients
// var broadcast = make(chan ProductUpdate)     // Broadcast channel for product updates
// var mutex = &sync.Mutex{}                    // Mutex for managing concurrent access

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		// Allow requests from the frontend (localhost:3000)
// 		return r.Header.Get("Origin") == "http://localhost:3000"
// 	},
// }

// // ProductUpdate struct to represent updates
// type ProductUpdate struct {
// 	ProductID string `json:"productId"`
// 	Quantity  int    `json:"quantity"`
// }

// // HandleConnections upgrades the HTTP connection to WebSocket
// func HandleConnections(w http.ResponseWriter, r *http.Request) {
// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Printf("WebSocket upgrade error: %v", err)
// 		http.Error(w, "Failed to upgrade to WebSocket", http.StatusBadRequest)
// 		return
// 	}
// 	defer ws.Close()

// 	// Add the new client
// 	mutex.Lock()
// 	clients[ws] = true
// 	mutex.Unlock()
// 	log.Println("New WebSocket connection established")

// 	// Keep the connection alive and listen for messages
// 	for {
// 		_, _, err := ws.ReadMessage()
// 		if err != nil {
// 			if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
// 				log.Printf("WebSocket client disconnected: %v", err)
// 			} else {
// 				log.Printf("WebSocket read error: %v", err)
// 			}
// 			// Remove the client on disconnect
// 			mutex.Lock()
// 			delete(clients, ws)
// 			mutex.Unlock()
// 			break
// 		}
// 	}
// }

// // HandleMessages listens for broadcast updates and sends them to all connected clients
// func HandleMessages() {
// 	for {
// 		// Wait for product updates
// 		update := <-broadcast
// 		mutex.Lock() // Lock when iterating over clients
// 		for client := range clients {
// 			err := client.WriteJSON(update)
// 			if err != nil {
// 				log.Printf("Error broadcasting message: %v", err)
// 				client.Close()
// 				delete(clients, client)
// 			}
// 		}
// 		mutex.Unlock()
// 	}
// }
