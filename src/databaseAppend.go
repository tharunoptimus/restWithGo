package src

func AddBooks(books []Book) []Book {
	books = append(books, Book{
		ID:    "1",
		Isbn:  "12341234",
		Title: "Book One",
		Author: &Author{
			Firstname: "meow", Lastname: "man",
		},
	})

	books = append(books, Book{
		ID:    "2",
		Isbn:  "67565",
		Title: "Book Two",
		Author: &Author{
			Firstname: "Thunder", Lastname: "Man",
		},
	})

	return books
}