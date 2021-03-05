package main

import (

	"github.com/gorilla/mux"
)

func MainRouter() *mux.Router{
	router := mux.NewRouter()
	AddSectionRoutes(router)
	return router
}

