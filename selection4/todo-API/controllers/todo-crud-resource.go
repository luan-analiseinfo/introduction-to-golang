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
			readAll(w, r)
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
}

func update(w http.ResponseWriter, r *http.Request) {

}

func readAll(w http.ResponseWriter, r *http.Request) {
	data, err := model.ReadAll()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(data)
	log.Println(r.URL)
}

func create(w http.ResponseWriter, r *http.Request) bool {
	data := views.PostRequest{}
	json.NewDecoder(r.Body).Decode(&data)
	if err := model.CreateTodo(&data); err != nil {
		w.Write([]byte("Some error"))
		return true
	}
	w.WriteHeader(http.StatusOK)
	return false
}
