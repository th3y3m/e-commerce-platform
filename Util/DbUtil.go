package Util

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

// ConnectToSQLServer connects to a SQL Server database
func ConnectToSQLServer() (*sql.DB, error) {
	connString := "server=(local);user id=sa;password=12345;database=ChatApplicationDb;encrypt=disable"

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
