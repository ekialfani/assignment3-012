package main

import (
	"assignment3-012/internal/router"
)

func main() {
	server := router.StartApp()
	server.Run(":8080")
}