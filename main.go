package main

import (
	"log"
	"net/http"
)

type ToDo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`

	Attributes map[string]any `json:"attributes"`
}

var todos = []ToDo{}

func newServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /todos/", listToDos)
	mux.HandleFunc("POST /todos/", addToDo)
	mux.HandleFunc("PATCH /todos/", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("DELETE /todos/", func(w http.ResponseWriter, r *http.Request) {})

	return mux
}

func main() {
	mux := newServer()

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalln(err.Error())
	}
}
