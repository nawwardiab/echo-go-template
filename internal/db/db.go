package db

import (
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

// ERRDBConnection – Returned when db connection fails
var ErrDBConnection = errors.New("database: connection failed")

// NewDB initializes and returns a new connection

func NewDB()(*pgx.Conn, error){
	// retrieve db connection string values from .env
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PWD")
	dbName := os.Getenv("DB_NAME")
	
	// DB_URL
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPwd, dbHost, dbPort, dbName)
	
	// initialize db connection
	config, parsingErr := pgx.ParseURI(connStr)
	conn, connectionErr := pgx.Connect(config)

	if parsingErr != nil {
		return nil, fmt.Errorf("invalid connection string: %w", parsingErr)
	} else if connectionErr != nil {
		return nil, fmt.Errorf("%w: %v", ErrDBConnection, connectionErr)
	} else {
		return conn, nil
	}
}