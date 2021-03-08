package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func MainRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/tasks", getTasks).Methods("POST")
	return router
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Task - POST")
}
