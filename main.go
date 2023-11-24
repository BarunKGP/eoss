package main

import (
	"eoss/pkg/handlers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// TODO: Implement client to avoid no timeout issues.
// TODO: Refer: https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779

func init() {
	// loads values from .env into system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	fmt.Println("Hello world!")

	router := http.HandlerFunc(handlers.Serve)
	port := os.Getenv("PORT")

	log.Panic(
		http.ListenAndServe(fmt.Sprintf(":%s", port), router),
	)
}
