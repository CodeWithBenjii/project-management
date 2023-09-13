package main

import (
	"log"

	"benjaminlee.dev/ProjectManagement/config"
	"benjaminlee.dev/ProjectManagement/db"
	"benjaminlee.dev/ProjectManagement/server"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config file: %v", err)
	}

	db.Initialize(cfg)
	defer db.Close()
	server.InitServer()
}
