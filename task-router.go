/*package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func AddTasksRoutes(router *mux.Router) {

	router.HandleFunc("/tasks", getTasks).Methods("POST")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)

}*/
