package main

import (
	pages "bargerbites/Pages"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// initEvents()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", pages.HomeLink)
	router.HandleFunc("/event", pages.CreateEvent).Methods("POST")
	router.HandleFunc("/events", pages.GetAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", pages.GetOneEvent).Methods("GET")
	router.HandleFunc("/events/{id}", pages.UpdateEvent).Methods("PATCH")
	router.HandleFunc("/events/{id}", pages.DeleteEvent).Methods("DELETE")
	log.Println("http://localhost:8080/")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
