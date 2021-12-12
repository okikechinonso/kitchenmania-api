package main

import (
	"kitchenmaniaapi/application/handlers"
	"kitchenmaniaapi/application/server"
	"kitchenmaniaapi/infrastructure/persistence"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("GIN_MODE")
	if env != "release" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("couldn't load env vars: %v", err)
		}
	}

	db := persistence.Database{}
	db.Init()

	s := server.Server{App: handlers.App{DB: db}}
	s.Start()
}
