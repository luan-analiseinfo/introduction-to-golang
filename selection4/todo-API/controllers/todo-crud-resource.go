package controllers

import (
	"../model"
	"../views"
	"encoding/json"
	"log"
	"net/http"
)

func crud() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			readTodo(w, r)
		case http.MethodPost:
			create(w, r)
		case http.MethodPut:
			update(w, r)
		case http.MethodDelete:
			deleteTodo(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]
	if err := model.DeleteTodo(name); err != nil {
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(struct {
		Status string `json:"status"`
	}{"Item deleted"})
	w.WriteHeader(http.StatusOK)
}

func update(w http.ResponseWriter, r *http.Request) {

}

func readTodo(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	data, err := model.ReadTodo(params)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(data)
	w.WriteHeader(http.StatusOK)
	log.Println(r.URL)
}

func create(w http.ResponseWriter, r *http.Request) bool {
	data := views.PostRequest{}
	json.NewDecoder(r.Body).Decode(&data)
	if err := model.CreateTodo(&data); err != nil {
		w.Write([]byte("Some error"))
		return true
	}
	w.WriteHeader(http.StatusCreated)
	return false
}
