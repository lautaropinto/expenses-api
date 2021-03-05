package main

import (
	"encoding/json"
	"net/http"
)

type section struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
}
type allsection []section

var sections = allsection{
	{
		ID:   1,
		Name: "sections",
	},
}

func (hc section) home(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(sections)
}
