package main

import (
	"database/sql"
	"log"

	"github.com/Jon-Castro856/Gator/internal/config"
	"github.com/Jon-Castro856/Gator/internal/database"
)

func initConfig(dbURL string) *config.State {
	jsonConfig, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error establishing database connection: %v", err)
	}

	return &config.State{
		Config: &jsonConfig,
		DB:     database.New(db),
	}
}
