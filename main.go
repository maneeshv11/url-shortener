package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"url-shortener/service"
	. "url-shortener/utils"
)

func init() {

	appConfig := LoadConfig()
	log.Printf("Initializing url shortner application on port: %s", appConfig.Port)

}

func main() {

	router := mux.NewRouter()

	// internal
	router.HandleFunc("/internal/queueUpdate", service.HandlePopulateQueue).Methods("GET")

	//external
	router.HandleFunc("/short", service.ShortenUrl).Methods("POST")
	router.HandleFunc("/{shortCode}", service.HandleOriginalUrlRequest)

	log.Fatal(http.ListenAndServe(":"+string(GetConfig().Port), router))

}
