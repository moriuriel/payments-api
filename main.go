package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/moriuriel/go-payments/infrastructure"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	infrastructure.NewHTTPServer().Start()
}
