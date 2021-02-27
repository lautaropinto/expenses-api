package main

import (
	"github.com/gorilla/mux"
)

func MainRouter() *mux.Router {
	router := mux.NewRouter()
	AddHomeRoutes(router)
	AddHelpRoutes(router)
	AddTasksRoutes(router)
	return router
}
