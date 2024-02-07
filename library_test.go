package model

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

// create a test setup for the library which will be used in the tests. it should initlaize the library and add some books to the library
func setupLibrary() Library {
	// Create a new instance of the library
	l := NewLibrary()

	// Add some books to the library
	book1 := Book{
		ISBN:  "1234567890",
		Title: "Book 1",
		Genre: "Fiction",
	}
	book2 := Book{
		ISBN:  "0987654321",
		Title: "Book 2",
		Genre: "Non-Fiction",
	}
	l.repository.AddBook(book1)
	l.repository.AddBook(book2)
}

// creat a teardown function for the library which will be used in the tests. it should remove all the books from the library
func teardownLibrary(l Library) {
	// Remove all books from the library
	l.repository.RemoveBook("1234567890")
	l.repository.RemoveBook("0987654321")
}

// create a utility fuction that can create a mock library for the tests
func createMockLibrary() Library {
	return &library{
		repository: &mockDatabase{},
	}
}

func TestLibrary_GetAllBooks(t *testing.T) {
	// Create a new instance of the library
	l := NewLibrary()

	// Add some books to the library
	book1 := Book{
		ISBN:   "1234567890",
		Title:  "Book 1",
		Author: "Author 1",
	}
	book2 := Book{
		ISBN:   "0987654321",
		Title:  "Book 2",
		Author: "Author 2",
	}
	l.repository.AddBook(book1)
	l.repository.AddBook(book2)

	// Call the GetAllBooks method
	books, err := l.GetAllBooks()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check if the returned books match the expected books
	expectedBooks := []Book{book1, book2}
	if !reflect.DeepEqual(books, expectedBooks) {
		t.Errorf("Expected books %v, got %v", expectedBooks, books)
	}
}

func TestLibrary_GetBook(t *testing.T) {
	// Create a new instance of the library
	l := NewLibrary()

	// Add a book to the library
	book := Book{
		ISBN:   "1234567890",
		Title:  "Book 1",
		Author: "Author 1",
	}
	l.repository.AddBook(book)

	// Test case 1: Get an existing book
	gotBook, err := l.GetBook("1234567890")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if gotBook != book {
		t.Errorf("Expected book %v, got %v", book, gotBook)
	}

	// Test case 2: Get a non-existing book
	_, err = l.GetBook("0987654321")
	if !errors.Is(err, ErrBookNotFound) {
		t.Errorf("Expected error %v, got %v", ErrBookNotFound, err)
	}
}

// Generate unit tests for GetAllFictionBooks
func TestGetAllFictionBooks(t *testing.T) {
	// Create a new instance of the library
	l := NewLibrary()

	// Add some books to the library
	book1 := Book{
		ISBN:  "1234567890",
		Title: "Book 1",
	}
	book2 := Book{
		ISBN:  "0987654321",
		Title: "Book 2",
	}
	l.repository.AddBook(book1)
	l.repository.AddBook(book2)

	// Call the GetAllFictionBooks method
	books := GetAllRecentFictionBooks(l)
	// Check if the returned books match the expected books
	expectedBooks := []Book{book1, book2}
	if !reflect.DeepEqual(books, expectedBooks) {
		t.Errorf("Expected books %v, got %v", expectedBooks, books)
	}
}

// Add a test case for GetAllFictionBooks where there are no fiction books in the library
func TestGetAllFictionBooksNoFictionBooks(t *testing.T) {
	// Create a new instance of the library
	l := NewLibrary()

	// Add some books to the library
	book1 := Book{
		ISBN:  "1234567890",
		Title: "Book 1",
	}
	book2 := Book{
		ISBN:  "0987654321",
		Title: "Book 2",
	}
	l.repository.AddBook(book1)
	l.repository.AddBook(book2)

	// Call the GetAllFictionBooks method
	books := GetAllRecentFictionBooks(l)
	// Check if the returned books match the expected books
	expectedBooks := []Book{}
	if !reflect.DeepEqual(books, expectedBooks) {
		t.Errorf("Expected books %v, got %v", expectedBooks, books)
	}
}

// Add a test case for GetAllFictionBooks where there are some fiction books in the library
func TestGetAllFictionBooksSomeFictionBooks(t *testing.T) {
	// Create a new instance of the library
	l := NewLibrary()

	// Add some books to the library
	book1 := Book{
		ISBN:      "1234567890",
		Title:     "Book 1",
		IsFiction: true,
	}
	book2 := Book{
		ISBN:      "0987654321",
		Title:     "Book 2",
		IsFiction: false,
	}
	l.repository.AddBook(book1)
	l.repository.AddBook(book2)

	// Call the GetAllFictionBooks method
	books := GetAllRecentFictionBooks(l)
	// Check if the returned books match the expected books
	expectedBooks := []Book{book1}
	if !reflect.DeepEqual(books, expectedBooks) {
		t.Errorf("Expected books %v, got %v", expectedBooks, books)
	}
}

// Add a test case for GetAllFictionBooks where GetAllLines returns an error
func TestGetAllFictionBooksError(t *testing.T) {
	// Create a new instance of the library
	l := NewLibrary()

	// Call the GetAllFictionBooks method
	books := GetAllRecentFictionBooks(l)
	// Check if the returned books match the expected books
	expectedBooks := []Book{}
	if !reflect.DeepEqual(books, expectedBooks) {
		t.Errorf("Expected books %v, got %v", expectedBooks, books)
	}
}

// Add a test case for GetAllFictionBooks where GetAllLines returns an error and GetAllFictionBooks returns errors.New("error reading the file")
func TestGetAllFictionBooksErrorReadingFile(t *testing.T) {
	// Create a new instance of the library
	l := NewLibrary()

	// Call the GetAllFictionBooks method
	_, err := GetAllRecentFictionBooks(l)
	// Check if the returned error matches the expected error
	expectedError := errors.New("error reading the file")
	if !errors.Is(err, expectedError) {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

// Add a  test case  for GetAllFictionBooks where parseBook returns an error and GetAllFictionBooks returns errors.New("Parsing error")
func TestGetAllFictionBooksParsingError(t *testing.T) {
	// Create a new instance of the library
	l := NewLibrary()

	// Call the GetAllFictionBooks method
	_, err := GetAllRecentFictionBooks(l)
	// Check if the returned error matches the expected error
	expectedError := errors.New("Parsing error")
	if !errors.Is(err, expectedError) {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

// Add a test case for GetAllRecentFictionBooks to assert that older books than 2020 are not returned
func TestGetAllRecentFictionBooksOldBooks(t *testing.T) {
	// Create a new instance of the library
	l := NewLibrary()

	// Add some books to the library
	book1 := Book{
		ISBN:      "1234567890",
		Title:     "Book 1",
		Genre:     "Fiction",
		Pages:     300,
		Published: 2014,
	}
	book2 := Book{
		ISBN:      "0987654321",
		Title:     "Book 2",
		Genre:     "Fiction",
		Pages:     300,
		Published: 2015,
	}
	book3 := Book{
		ISBN:      "0987654321",
		Title:     "Book 3",
		Genre:     "Fiction",
		Pages:     300,
		Published: 2016,
	}
	l.repository.AddBook(book1)
	l.repository.AddBook(book2)
	l.repository.AddBook(book3)

	// Call the GetAllRecentFictionBooks method
	books := GetAllRecentFictionBooks(l)
	// Check if the returned books match the expected books
	expectedBooks := []Book{book2, book3}
	if !reflect.DeepEqual(books, expectedBooks) {
		t.Errorf("Expected books %v, got %v", expectedBooks, books)
	}
}

// write a testcase to test the highlighted text in this method:book.Pages > 300 && book.Published.Year() > 2020
func TestGetAllRecentFictionBooks(t *testing.T) {
	// Create a mock library
	mockLibrary := createMockLibrary()

	// Mock the GetAllLines method to return some lines
	mockLibrary.repository.(*mockDatabase).GetAllLinesFunc = func() ([]string, error) {
		return []string{
			"ISBN:1234567890,Title:Book 1,Genre:Fiction,Pages:400,Published:2021-01-01",
			"ISBN:0987654321,Title:Book 2,Genre:Fiction,Pages:200,Published:2022-02-02",
			"ISBN:9876543210,Title:Book 3,Genre:Non-Fiction,Pages:500,Published:2023-03-03",
		}, nil
	}

	// Call the GetAllRecentFictionBooks method
	fictionBooks := GetAllRecentFictionBooks(mockLibrary)

	// Check if the returned books match the expected books
	expectedBooks := []Book{
		{
			ISBN:      "1234567890",
			Title:     "Book 1",
			Genre:     "Fiction",
			Pages:     400,
			Published: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	if !reflect.DeepEqual(fictionBooks, expectedBooks) {
		t.Errorf("Expected books %v, got %v", expectedBooks, fictionBooks)
	}
}

// write a testcase to test the highlighted text in this method:ratingService.GetBookRating(&book) > 3
func TestGetAllRecentFictionBooksRating(t *testing.T) {
	// Create a mock library
	mockLibrary := createMockLibrary()

	// Mock the GetAllLines method to return some lines
	mockLibrary.repository.(*mockDatabase).GetAllLinesFunc = func() ([]string, error) {
		return []string{
			"ISBN:1234567890,Title:Book 1,Genre:Fiction,Pages:400,Published:2021-01-01",
			"ISBN:0987654321,Title:Book 2,Genre:Fiction,Pages:200,Published:2022-02-02",
			"ISBN:9876543210,Title:Book 3,Genre:Non-Fiction,Pages:500,Published:2023-03-03",
		}, nil
	}

	// Mock the GetBookRating method to return a rating of 4
	mockLibrary.ratingService.(*mockRatingService).GetBookRatingFunc = func(book *Book) int {
		return 4
	}

	// Call the GetAllRecentFictionBooks method
	fictionBooks := GetAllRecentFictionBooks(mockLibrary)

	// Check if the returned books match the expected books
	expectedBooks := []Book{
		{
			ISBN:      "1234567890",
			Title:     "Book 1",
			Genre:     "Fiction",
			Pages:     400,
			Published: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			ISBN:      "0987654321",
			Title:     "Book 2",
			Genre:     "Fiction",
			Pages:     200,
			Published: time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC),
		},
	}
	if !reflect.DeepEqual(fictionBooks, expectedBooks) {
		t.Errorf("Expected books %v, got %v", expectedBooks, fictionBooks)
	}
}
