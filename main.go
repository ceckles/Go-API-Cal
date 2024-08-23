package main

import (
	"github.com/gin-gonic/gin"

	//"fmt"
	"log"
	"net/http"
)

type SumRequest struct {
	Sum1 int `json:"sum1" bindings.Required`
	Sum2 int `json:"sum2" bindings.Required`
}
type SubRequest struct {
	Number1 int `json:"number1" bindings.Required`
	Number2 int `json:"number2" bindings.Required`
}

func main() {
	// Create a new Gin router with default logger and recovery middleware
	r := gin.Default()

	// Define a route with a path variable
	r.GET("/hello/:name", func(c *gin.Context) {
		// Get the path variable
		name := c.Param("name")

		// Respond with a message
		c.String(200, "Hello, %s", name)
	})

	// Define the /sum endpoint
	r.POST("/sum", func(c *gin.Context) {
		var req SumRequest
	
		// Bind and validate JSON input
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Calculate the sum
		result := req.Sum1 + req.Sum2
	
		// Respond with the sum
		c.JSON(http.StatusOK, gin.H{"sum": result})
	})

	// Define the /subtract endpoint
	r.POST("/subtract", func(c *gin.Context){
		var req SubRequest

		// Bind and validate JSON input
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Calculate the Subtraction
		result := req.Number1 - req.Number2
		
		// Respond with the subtraction
		c.JSON(http.StatusOK, gin.H{"subtraction": result})
	})

	// Log and start the server
	
	log.Println("Starting Server on port 3000")
	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
