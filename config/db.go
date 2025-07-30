package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/microsoft/go-mssqldb"
)

// NewConnection opens a SQL Server connection using environment variables:
// DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME
func NewConnection() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		user, pass, host, port, name)
	dbClient, err := sql.Open("sqlserver", dsn)
	if err != nil {
		return nil, err
	}

	// Verify connection
	if err := dbClient.Ping(); err != nil {
		return nil, err
	}

	// Connection pool settings
	dbClient.SetMaxOpenConns(25)
	dbClient.SetMaxIdleConns(25)
	dbClient.SetConnMaxLifetime(5 * time.Minute)

	return dbClient, nil
}
