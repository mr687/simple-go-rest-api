package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gitlab.com/mr687/privy-be-test-go/config"
	"gitlab.com/mr687/privy-be-test-go/controller"
)

// Main function
func main() {
	// Load env file
	LoadEnvironment()

	// Start GO server
	server := controller.Server{}
	server.Initialize(config.DatabaseConfig{
		DBHost: os.Getenv("DB_PG_HOST"),
		DBPort: os.Getenv("DB_PG_PORT"),
		DBUser: os.Getenv("DB_PG_USER"),
		DBPass: os.Getenv("DB_PG_PASSWORD"),
		DBName: os.Getenv("DB_PG_DATABASE"),
	})
	server.StartServer(config.ServerConfig{
		Port: os.Getenv("PORT"),
	})
}

func LoadEnvironment() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("An error occured when getting .env file %v", err)
	}
}
