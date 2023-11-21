package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main () {
	// fmt.Println("Hello world!")
	dotenv := loadEnvVars("CLIENT_ID")
	fmt.Printf("godotenv: %s=%s \n", "CLIENT_ID", dotenv)
}

func loadEnvVars(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

