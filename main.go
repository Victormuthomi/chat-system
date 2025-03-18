package main

import (
	"fmt"
	"log"

	"chat-system/config"
	"chat-system/routes"
	"chat-system/ws"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration.
	cfg := config.LoadConfig()

	// Initialize Gin router.
	router := gin.Default()

	// Initialize WebSocket Hub.
	hub := ws.NewHub()
	go hub.Run()

	// Register routes.
	routes.RegisterRoutes(router, hub)

	// Start the server.
	port := cfg.Port
	fmt.Println("🚀 Server is running on port", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Server failed:", err)
	}
}

