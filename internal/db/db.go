package db

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx"
)

// ERRDBConnection – Returned when db connection fails
var ErrDBConnection = errors.New("database: connection failed")

// NewDB initializes and returns a new connection

func NewDB(connStr string)(*pgx.Conn, error){
	config, parsingErr := pgx.ParseURI(connStr)
	if parsingErr != nil {
		return nil, fmt.Errorf("invalid connection string: %w", parsingErr)
	}

	conn, connectionErr := pgx.Connect(config)
	if connectionErr != nil {
		return nil, fmt.Errorf("%w: %v", ErrDBConnection, connectionErr)
	}

	return conn, nil
}