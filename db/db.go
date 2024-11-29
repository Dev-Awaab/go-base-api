package db

import (
	"database/sql"

	"github.com/Dev-Awaab/go-base-api/pkg/logger"
	_ "github.com/lib/pq"
)

func  InitDB(dbSource string) *sql.DB  {
	db, err := sql.Open("postgres", dbSource)

	if err != nil {
		logger.Error("Failed to connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		logger.Error("Database connection error: %v", err)
	}
	logger.Error("Connected to the database successfully!")
	return db
}