CRUD REST API with Gorilla Mux
## Basic

### Initialize the project
`go mod init restWithGo`

### Install the dependencies
`go get -u github.com/gorilla/mux`

### Run the application
`go run restWithGo.go`

### Build
`go build restWithGo.go`

## Info
- A Simple CRUD REST API to create, read, update and delete books.
- The API is built with the Gorilla Mux library.
- The Database has been mocked with a slice in the memory

- `structures.go` holds the structs for the data objects for Books and Author
- `databaseAppend.go` will append the data to the slice in memory
- `crud.go` contains the CRUD functions that the router will perform

- Production ready? No! Template worthy? Yes!

Enjoy 🎉