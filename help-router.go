package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func AddHelpRoutes(router *mux.Router) {
	router.HandleFunc("/help", help).Methods("GET")
	router.HandleFunc("/help", newHelp).Methods("POST")

}

func help(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("HELP - GET")
}

func newHelp(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("new help - POST")
}
