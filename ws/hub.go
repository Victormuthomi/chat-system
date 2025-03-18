package ws

import (
	"fmt"
)

// Hub maintains active clients and broadcasts messages.
type Hub struct {
	Clients    map[*Client]bool // Active clients
	Broadcast  chan []byte      // Channel for broadcasting messages
	Register   chan *Client     // Channel for new client registrations
	Unregister chan *Client     // Channel for client disconnections
}

// NewHub creates and returns a new Hub.
func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

// Run starts the hub loop to manage clients.
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			fmt.Println("✅ New client connected")
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
				fmt.Println("❌ Client disconnected")
			}
		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}

