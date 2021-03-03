package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func MainRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/sections", getSections).Methods("GET")
	return router
}


func getSections(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Sections - GET")
}