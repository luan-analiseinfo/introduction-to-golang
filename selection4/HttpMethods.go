package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("requested received")
		log.Println(request.Method)
		writer.Write([]byte("Hello Word"))
	})
	mux.HandleFunc("/go", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("requested received")
		log.Println(request.Method)
		writer.Write([]byte("Hello Word"))
	})
	http.ListenAndServe("localhost:3000", mux)
}
