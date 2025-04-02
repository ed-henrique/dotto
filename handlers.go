package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func listToDos(w http.ResponseWriter, r *http.Request) {
	var res string

	for _, todo := range todos {
		res = fmt.Sprintf("%s%s\n", res, todo.Name)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, res)
}

func addToDo(w http.ResponseWriter, r *http.Request) {
	var todo ToDo

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&todo); err != nil {
		log.Println(err.Error())
	}

	w.WriteHeader(http.StatusCreated)
}
