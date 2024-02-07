package model

import "errors"

/*
	create a new interface called library.

it should have methods called GetBook, DeleteBook, and AddBook.
it should use the book from the book package.
*/
type Library interface {
	GetBook(isbn string) (Book, error)
	DeleteBook(isbn string) error
	AddBook(b Book) error
}

type library struct {
	repository Database // Add a field called
}

func NewLibrary() Library {
	return &library{
		repository: NewFileDatabase("books.txt"),
	}
}

func (l *library) GetBook(isbn string) (Book, error) {
	//if the string is less the 4 characters short return error
	if len(isbn) < 4 {
		return Book{}, errors.New("isbn must be at least 4 characters long")
	}
	//if the string is more than 12 characters long return error
	if len(isbn) > 12 {
		return Book{}, errors.New("isbn must be at most 12 characters long")
	}
	//get the book from the database
	book, err := l.repository.GetLine(isbn)

	return nil, nil
}

func (l *library) DeleteBook(isbn string) error {
	delete(l.books, isbn)
	return nil
}

func (l *library) AddBook(b Book) error {
	l.books[b.ISBN] = b
	return nil
}

func (l *library) GetAllBooks() ([]Book, error) {
	// Implement the logic to get all books from the database
	// Get all the books from the database
	books, err := l.repository.GetAllLines()
	if err != nil {
		return nil, err
	}

	// Convert the lines to books
	var allBooks []Book
	for _, line := range books {
		book, err := parseBook(line)
		if err != nil {
			return nil, err
		}
		allBooks = append(allBooks, book)
	}

	return allBooks, nil
}

func GetAllRecentFictionBooks(l library) []Book {
	var fictionBooks []Book

	books, err := l.repository.GetAllLines()
	if err != nil {
		return nil, errors.New("error reading the file")
	}

	// Convert the lines to books
	var allBooks []Book
	for _, line := range books {
		book, err := parseBook(line)
		if err != nil {
			return nil, errors.New("Parsing error")
		}
		allBooks = append(allBooks, book)
	}

	for _, book := range allBooks {
		if book.Genre == "Fiction" {
			if book.Pages > 300 && book.Published.Year() > 2020 {
				fictionBooks = append(fictionBooks, book)
			}
		}
	}
	return fictionBooks
}
