package main

import (
	"assignment3-012/internal/router"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}
	
	server := router.StartApp()
	server.Run(":8080")
}