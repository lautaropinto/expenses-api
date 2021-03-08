package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func MainRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/tasks", getTasks).Methods("POST")
	router.HandleFunc("/home", resolveHome).Methods("GET")
	return router
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Task - POST")
}

func resolveHome(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("GET HOME PAPURRI")
}
