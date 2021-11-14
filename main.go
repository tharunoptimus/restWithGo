package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Initialize books variable as a slide Book struct
var books []Book



func main() {

	var dbConnectionString string = goDotEnvVariable("MONGO_URI")
	var port string = goDotEnvVariable("PORT")

	client, ctx, cancel, err := connect(dbConnectionString)
    if err != nil {
        panic(err)
    }
     
    // Release resource when the main
    // function is returned.
    defer close(client, ctx, cancel)
     
    // Ping mongoDB with Ping method
	ping(client, ctx)

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

	// Start the server
	fmt.Println("Server listening on port "+ port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
