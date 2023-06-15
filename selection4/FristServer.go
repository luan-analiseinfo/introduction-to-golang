package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("received")
		writer.Write([]byte("Hello Word"))
	})
	http.ListenAndServe("localhost:3000", mux)
}
