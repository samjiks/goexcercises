package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/samjiks/lesson2/webservicetask/handlers"
)

func main() {
	port := flag.String("port", "8000", "Server port")
	flag.Parse()
    	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/", handlers.HelloHandler)
	router.HandleFunc("/developers", handlers.DevelopersHandler).Methods("GET", "POST")
	router.HandleFunc("/developers/{id}", handlers.DevelopersHandler).Methods("GET")

	// Starts the web server
	log.Fatal(http.ListenAndServe(":"+*port, router))
}
