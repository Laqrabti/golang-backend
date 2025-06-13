package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	router := setupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Server running on port %s", port)
	log.Fatal(router.Run(":" + port))
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	
	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Routes
	r.GET("/health", healthCheck)
	r.GET("/ping", pingHandler)
	
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/status", apiStatus)
		apiGroup.POST("/echo", echoHandler)
	}

	return r
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "UP",
		"version": "1.0.0",
	})
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"time":    time.Now().Format(time.RFC3339),
	})
}

func apiStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"service": "Golang Backend API",
		"status":  "operational",
		"uptime":  time.Since(startTime).String(),
	})
}

func echoHandler(c *gin.Context) {
	var request struct {
		Message string 
	}
	
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"echo":    request.Message,
		"reverse": reverseString(request.Message),
	})
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

var startTime = time.Now()
