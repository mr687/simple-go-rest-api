package config

import (
	"fmt"
	"log"
	"os"

	"gitlab.com/mr687/privy-be-test-go/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() *gorm.DB {
	fmt.Println("Connecting to database...")

	// Define database config
	db_host := os.Getenv("DB_PG_HOST")
	db_port := os.Getenv("DB_PG_PORT")
	db_user := os.Getenv("DB_PG_USER")
	db_pass := os.Getenv("DB_PG_PASSWORD")
	db_name := os.Getenv("DB_PG_DATABASE")

	uri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", db_user, db_pass, db_host, db_port, db_name)

	// Open a new connection to database
	db, err := gorm.Open(postgres.Open(uri))
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func DBAutoMigrate(db *gorm.DB) {
	fmt.Println("Migrating database...")
	db.AutoMigrate(
		entity.User{},
		entity.UserBalance{},
		entity.BankBalance{},
		entity.UserBalanceHistory{},
		entity.BankBalanceHistory{},
	)
}

func CloseConnection(db *gorm.DB) {
	fmt.Println("Closing database connection...")
	sql, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sql.Close()
}
