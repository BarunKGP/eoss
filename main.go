package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"eoss/pkg/handlers"

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

	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/login/github/", handlers.GithubLoginHandler)
	http.HandleFunc("/login/github/callback", handlers.GithubCallbackHandler)
	http.HandleFunc("/loggedin", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoggedinHandler(w, r, "")
	})

	port := os.Getenv("PORT")
	log.Panic(
		http.ListenAndServe(":"+string(port), nil),
	)
}


// func loadEnvVars(key string) string {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}
// 	return os.Getenv(key)
// }
