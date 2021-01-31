package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server init!")
	router := MainRouter()

	router.Use(defaultHeadersMiddleware)

	http.ListenAndServe(":8000", router)
}
