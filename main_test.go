package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddToDo(t *testing.T) {
	original := ToDo{
		ID:         0,
		Name:       "a",
		Done:       false,
		Attributes: nil,
	}

	var body bytes.Buffer
	encoder := json.NewEncoder(&body)
	_ = encoder.Encode(&original)

	t.Run("valid todo", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/todos", &body)
		w := httptest.NewRecorder()
		addToDo(w, req)

		res := w.Result()
		defer res.Body.Close()

		if http.StatusCreated != res.StatusCode {
			t.Errorf("expected %d got %d", http.StatusCreated, res.StatusCode)
		}
	})
}
