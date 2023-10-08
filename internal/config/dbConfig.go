package config

import (
	"fmt"
	"os"
)

func DBConfig() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	return fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s", host, user, password, port, dbname, sslmode)
}