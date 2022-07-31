package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.com/mr687/privy-be-test-go/config"
	"gitlab.com/mr687/privy-be-test-go/controller"
	"gitlab.com/mr687/privy-be-test-go/repository"
	"gitlab.com/mr687/privy-be-test-go/service"
)

// Main function
func main() {
	// Load env file
	LoadEnvironment()

	dbConn := config.NewConnection()

	// Close database connection when program exits
	defer config.CloseConnection(dbConn)

	// Run auto migrate
	config.DBAutoMigrate(dbConn)

	router := gin.Default()

	// Factory for repository
	userRepository := repository.NewUserRepository(dbConn)

	// Factory for service
	authService := service.NewAuthService(userRepository)
	jwtService := service.NewJwtService()

	// Factory for controller
	authController := controller.NewAuthController(authService, jwtService)

	// Register routes
	apiV1 := router.Group("/api/v1")
	{
		authApi := apiV1.Group("/auth")
		{
			authApi.POST("/register", authController.Register)
			authApi.POST("/login", authController.Login)
		}
	}
	router.NoRoute(func(ctx *gin.Context) {
		ctx.AbortWithStatus(404)
	})

	// Define port or default port if None
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
