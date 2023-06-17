package main

import (
	"./controllers"
	"./model"
	"log"
	"net/http"
)

func main() {
	mux := controllers.Register()
	db := model.Connect()
	defer db.Close()
	log.Println("Serving...")
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
