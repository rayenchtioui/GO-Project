package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	var errOpen error
	db, errOpen = sql.Open("mysql", dsn)
	if errOpen != nil {
		log.Fatalf("Error opening database: %v", errOpen)
	}

	errPing := db.Ping()
	if errPing != nil {
		db.Close()
		log.Fatalf("Error connecting to database: %v", errPing)
	}

	fmt.Println("Successfully connected to the database")
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error closing database: %v", err)
		}
		fmt.Println("Database connection closed successfully")
	}
}

func DeleteDB() error {
	dbName := os.Getenv("DB_NAME")
	_, err := db.Exec("DROP DATABASE IF EXISTS " + dbName)
	if err != nil {
		return fmt.Errorf("failed to delete database: %v", err)
	}
	return nil
}
