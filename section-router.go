package main

import (

	"github.com/gorilla/mux"
)

func AddSectionRoutes(router *mux.Router) {
	Controller := section{}
	router.HandleFunc("/section", Controller.home).Methods("GET")
}
