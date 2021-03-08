package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Server init!")
	router := MainRouter()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	host := os.Getenv("PORT")

	if host == "" {
		host = "3000"
	}

	http.ListenAndServe(":" + host, router)

}
