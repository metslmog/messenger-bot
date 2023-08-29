package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"messenger-bot/lib"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/webhook", lib.VerificationEndpoint).Methods("GET")
	r.HandleFunc("/webhook", lib.MessagesEndpoint).Methods("POST")
	if err := http.ListenAndServe("0.0.0.0:8080", r); err != nil {
		log.Fatal(err)
	}
}