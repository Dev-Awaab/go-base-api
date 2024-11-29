package main

import (
	"fmt"

	"github.com/Dev-Awaab/go-base-api/config"
	"github.com/Dev-Awaab/go-base-api/db"
	"github.com/Dev-Awaab/go-base-api/pkg/logger"
	"github.com/Dev-Awaab/go-base-api/pkg/router"
)


func main(){
	// Initialize the logger
	_, err := logger.InitLogger(logger.InfoLevel, "stdout")
	if err != nil {
	 panic(err)
	}
	// Load configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		logger.Error("Failed to load configuration: %v", err)
	}

	// Initialize the database 
	dbConn := db.InitDB(cfg.DBSource)

	fmt.Println("ServerAddress",cfg)

	router.SetupRoutes(dbConn, cfg)
	router.Start(cfg.ServerAddress)
	logger.Info("Starting server on %s", cfg.ServerAddress)
}