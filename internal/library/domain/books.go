package domain

import "fmt"

type BookStatus string

const (
	BookStatusToDo     BookStatus = "TO_DO"
	BookStatusBooked   BookStatus = "BOOKED"
	BookStatusAvaiable BookStatus = "AVAILABLE"
	BookStatusDeleted  BookStatus = "DELETED"
)

type Book struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	IsBorrowed bool   `json:"is_borrowed"`
}

type BooksRepository interface {
	AddBook(input AddBookInput) error
	BorrowBook(input BorrowBookInput) error
	ReturnBook(input BorrowBookInput) error
	GetBook(input GetBookInput) (Book, error)
	PutBook(input PutBookInput) (Book, error)
	DeleteBook(input DeleteBookInput) error
	ListBooks() []Book
}

type (
	AddBookInput struct {
		BookID string
		Title  string `json:"title"`
		Author string `json:"author"`
	}
	BorrowBookInput struct {
		BookID string
	}
	GetBookInput struct {
		BookID string
	}
	PutBookInput struct {
		BookID string
		Title  string `json:"title"`
		Author string `json:"author"`
	}
	DeleteBookInput struct {
		BookID string
	}
)

var (
	ErrBookCannotCreate       = fmt.Errorf("book cannot be created")
	ErrBookInvalidInputTitle  = fmt.Errorf("book title is invalid")
	ErrBookInvalidInputAuthor = fmt.Errorf("book author is invalid")
	ErrBookAlreadyBorrowed    = fmt.Errorf("book is already borrowed")
	ErrBookNotBorrowed        = fmt.Errorf("book is not borrowed")
	ErrBookNotFound           = fmt.Errorf("book not found")
)
