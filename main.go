package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

func main() { 
	fmt.Println("favio c ka come")
	router := mux.NewRouter()

	router.HandleFunc("/texto", textHandler).Methods("GET")

	http.ListenAndServe(":3000", router)
}

func textHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Text - GET")
}
