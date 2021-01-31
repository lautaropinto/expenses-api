package main

import (
	"encoding/json"
	"net/http"
)

type HomeController struct{}

func (hc HomeController) home(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Successx")
}
