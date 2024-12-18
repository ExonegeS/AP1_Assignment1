package service

import "github.com/ExonegeS/AP1_Assignment1/internal/library/domain"

type Service interface {
	AddBook(input AddBookInput) (AddBookResponse, error)
	GetBook(input GetBookInput) (GetBookResponse, error)
	BorrowBook(input BorrowBookInput) (BorrowBookResponse, error)
	ReturnBook(input BorrowBookInput) (BorrowBookResponse, error)
	PutBook(input PutBookInput) (PutBookResponse, error)
	DeleteBook(input DeleteBookInput) (DeleteBookResponse, error)
	ListBooks() ListBooksResponse
}

type (
	AddBookInput struct {
		Title  string `json:"title"`
		Author string `json:"author"`
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
	PutBookInput struct {
		BookID string
		Title  string `json:"title"`
		Author string `json:"author"`
	}
	PutBookResponse struct {
		Book domain.Book `json:"book"`
	}
)

type (
	DeleteBookInput struct {
		BookID string
	}
	DeleteBookResponse struct {
		Status domain.BookStatus `json:"book_status"`
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
