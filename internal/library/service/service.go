package service

import "github.com/ExonegeS/AP1_Assignment1/internal/library/domain"

type Service interface {
	AddBook(input AddBookInput) (AddBookResponse, error)
	GetBook(input GetBookInput) (GetBookResponse, error)
	BorrowBook(input BorrowBookInput) (BorrowBookResponse, error)
	ReturnBook(input BorrowBookInput) (BorrowBookResponse, error)
	ListBooks() ListBooksResponse
}

type (
	AddBookInput struct {
		Title  string
		Author string
	}
	AddBookResponse struct {
		Book domain.Book `json:"book"`
	}
)

type (
	GetBookInput struct {
		BookID string
	}
	GetBookResponse struct {
		Book domain.Book `json:"book"`
	}
)

type (
	BorrowBookInput struct {
		BookID string
	}
	BorrowBookResponse struct {
		Status domain.BookStatus `json:"book_status"`
	}
)

type (
	ListBooksResponse struct {
		Books []domain.Book `json:"books"`
	}
)

type service struct {
	books domain.BooksRepository
}

var _ Service = (*service)(nil)

func NewService(books domain.BooksRepository) Service {
	return &service{
		books: books,
	}
}