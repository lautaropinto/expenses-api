package main

import (
    "github.com/gorilla/mux"
)

func AddHomeRoutes(router *mux.Router) {
	Controller := HomeController{}
    router.HandleFunc("/home", Controller.home).Methods("GET")
}