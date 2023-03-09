package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ymho/qnin-sp/handlers"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.DefaultHandler).Methods(http.MethodGet)
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/saml", handlers.PostSAMLHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
