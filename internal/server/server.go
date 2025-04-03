package server

import (
	"io/fs"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/ed-henrique/dotto/internal/models"
	"github.com/ed-henrique/dotto/internal/render"
)

type Server struct {
	mux   *http.ServeMux
	todos []models.ToDo
	views fs.FS
}

func New(views fs.FS, todos []models.ToDo) *Server {
	return &Server{
		mux:   http.NewServeMux(),
		todos: todos,
		views: views,
	}
}

func (s *Server) Routes() {
	renderer := render.NewRenderer(s.views)

	s.mux.HandleFunc("GET /todos/test", func(w http.ResponseWriter, r *http.Request) {
		t, err := renderer.Render("views/components/base.html", "views/components/todo.html")
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(500)
			return
		}

		if err := t.Execute(w, s.todos); err != nil {
			log.Println(err.Error())
			w.WriteHeader(500)
			return
		}
	})

	s.mux.HandleFunc("GET /todos/", s.listToDos)
	s.mux.HandleFunc("POST /todos/", s.addToDo)
	s.mux.HandleFunc("PATCH /todos/", func(w http.ResponseWriter, r *http.Request) {})
	s.mux.HandleFunc("DELETE /todos/", func(w http.ResponseWriter, r *http.Request) {})
}

func (s *Server) Run() {
	if err := http.ListenAndServe(":8080", s.mux); err != nil {
		slog.Error("could not start server", slog.String("err", err.Error()))
		os.Exit(1)
	}
}
