package Util

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

// ConnectToSQLServer connects to a SQL Server database
func ConnectToSQLServer() (*sql.DB, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	databaseURL := os.Getenv("CONNECTION_STRING")

	connString := databaseURL

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to SQL Server!")
	return db, nil
}
