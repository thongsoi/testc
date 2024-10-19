package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		return err
	}

	// Get the connection string from the environment
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Println("DATABASE_URL is not set in the environment")
		return err
	}

	// Initialize the database connection
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Error opening database connection: %v", err)
		return err
	}

	// Set connection pool settings
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(time.Hour)

	// Ping database to check connection
	if err = DB.Ping(); err != nil {
		log.Printf("Error pinging database: %v", err)
		return err
	}
	log.Println("Database connection established")
	return nil
}

func CloseDB() error {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Printf("Error closing database connection: %v", err)
			return err
		}
		log.Println("Database connection closed")
	}
	return nil
}

// GetDB returns the existing DB instance
func GetDB() *sql.DB {
	if DB == nil {
		log.Fatal("Database connection is not initialized")
	}
	return DB
}
