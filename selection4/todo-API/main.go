package main

import (
	"./controllers"
	"./model"
	"net/http"
)

func main() {
	mux := controllers.Register()
	db := model.Connect()
	defer db.Close()
	http.ListenAndServe("localhost:3000", mux)
}
