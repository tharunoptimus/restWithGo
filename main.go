package main

import (
	"fmt"
	"log"
	"net/http"

	s "github.com/tharunoptimus/restWithGo/src"

	"github.com/gorilla/mux"
)

// Initialize books variable as a slide Book struct



func main() {

	var dbConnectionString string = s.GoDotEnvVariable("MONGO_URI")
	var port string = s.GoDotEnvVariable("PORT")

	client, ctx, cancel, err := s.Connect(dbConnectionString)
    if err != nil {
        panic(err)
    }
     
    // Release resource when the main
    // function is returned.
    defer s.Close(client, ctx, cancel)
     
    // Ping mongoDB with Ping method
	s.Ping(client, ctx)

	// Init the Mux router
	r := mux.NewRouter()

	// Mock Data -- @todo - implement a DB

	// Route handlers / Endpoints
	r.HandleFunc("/api/books", s.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", s.GetBook).Methods("GET")
	r.HandleFunc("/api/books", s.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", s.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", s.DeleteBook).Methods("DELETE")

	// Start the server
	fmt.Println("Server listening on port "+ port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
