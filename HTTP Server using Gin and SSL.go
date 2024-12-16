package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router instance
	r := gin.Default()

	// Define a simple GET endpoint to return a greeting
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, secure world!",
		})
	})

	// Define a POST endpoint for user login
	r.POST("/login", func(c *gin.Context) {
		var json struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		// Bind the JSON data to the struct
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		// Respond back with a success message
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Welcome, %s!", json.Username),
		})
	})

	// Start the HTTP server with SSL certificates
	if err := r.RunTLS(":8080", "cert.pem", "key.pem"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
