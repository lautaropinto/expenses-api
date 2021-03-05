package main

import (
	"fmt"
	"net/http"
	"os"

)

func main() {
	fmt.Println("Server init!")
	router := MainRouter()
	router.Use(defaultHeadersMiddleware)

	host := os.Getenv("DB_PORT")

	if host == "" {
		host = "localhost:3000"
	}

	http.ListenAndServe(host, router)

}
