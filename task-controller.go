package main

type task struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
}
type alltasks []task

var tasks = alltasks{
	{
		ID:   1,
		Name: "Fede",
	},
}
