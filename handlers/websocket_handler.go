package handlers

import (
	"net/http"

	"chat-system/ws"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Upgrader upgrades HTTP connections to WebSocket connections.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections (adjust in production)
	},
}

// WebSocketHandler upgrades HTTP requests to WebSocket and registers the client.
func WebSocketHandler(hub *ws.Hub, c *gin.Context) {
	// Upgrade HTTP to WebSocket using Gin's context.
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "WebSocket upgrade failed"})
		return
	}

	// Create a new client using the exported constructor from ws.Client.
	client := ws.NewClient(hub, conn)

	// Register the client with the hub.
	hub.Register <- client

	// Start routines to handle messages.
	go client.ReadMessages()
	go client.WriteMessages()
}

