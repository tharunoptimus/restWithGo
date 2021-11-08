package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Initialize books variable as a slide Book struct
var books []Book



func main() {
	// Init the Mux router
	r := mux.NewRouter()

	// Mock Data -- @todo - implement a DB
	books = addBooks(books)

	// Route handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
