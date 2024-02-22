package main

import (
	"logging-middleware/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Register your logging middleware
	r.Use(middleware.LoggingMiddleware(), middleware.ResponseLoggingMiddleware())

	// Define a simple handler for the `/example` endpoint
	r.POST("/example", func(c *gin.Context) {
		var requestJSON map[string]interface{}
		if err := c.BindJSON(&requestJSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Example processing and response
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": requestJSON})
	})

	// Run the server
	r.Run() // By default, it listens on :8080
}
