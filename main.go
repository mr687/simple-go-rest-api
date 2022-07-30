package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Main function
func main() {
	// Load env file
	LoadEnvironment()

	router := gin.Default()

	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Hello World!",
			})
		})
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.AbortWithStatus(404)
	})

	// Define port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	router.Run(":" + port)
}

// Load env file
func LoadEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}
