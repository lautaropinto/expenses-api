package main

type section struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
}
type allsection []task

var sections = allsection{
	{
		ID:   1,
		Name: "sections",
	},
}
