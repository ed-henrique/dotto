package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ed-henrique/dotto/internal/models"
)

func (s *Server) listToDos(w http.ResponseWriter, r *http.Request) {
	var res string

	for _, todo := range s.todos {
		res = fmt.Sprintf("%s%s\n", res, todo.Name)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, res)
}

func (s *Server) addToDo(w http.ResponseWriter, r *http.Request) {
	var todo models.ToDo

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&todo); err != nil {
		log.Println(err.Error())
	}

	w.WriteHeader(http.StatusCreated)
}
