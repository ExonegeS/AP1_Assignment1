package adapters

import (
	"github.com/ExonegeS/AP1_Assignment1/internal/library/domain"
)

type BooksRepositoryStruct struct {
	data map[string]domain.Book
}

var _ domain.BooksRepository = (*BooksRepositoryStruct)(nil)

func NewBooksRepositoryStruct() *BooksRepositoryStruct {
	return &BooksRepositoryStruct{
		data: make(map[string]domain.Book),
	}
}

func (b *BooksRepositoryStruct) AddBook(input domain.AddBookInput) error {
	// Add a new book to the repository

	newBook := domain.Book{
		ID:         input.BookID,
		Title:      input.Title,
		Author:     input.Author,
		IsBorrowed: false,
	}

	if _, ok := b.data[input.BookID]; ok {
		return domain.ErrBookCannotCreate
	}

	b.data[input.BookID] = newBook
	return nil
}

func (b *BooksRepositoryStruct) GetBook(input domain.GetBookInput) (domain.Book, error) {
	// Get a book from the repository

	bookID := input.BookID

	book, ok := b.data[bookID]
	if !ok {
		return domain.Book{}, domain.ErrBookNotFound
	}

	return book, nil
}

func (b *BooksRepositoryStruct) BorrowBook(input domain.BorrowBookInput) error {
	// Borrow a book from the repository

	bookID := input.BookID

	book, ok := b.data[bookID]
	if !ok {
		return domain.ErrBookNotFound
	}

	if book.IsBorrowed {
		return domain.ErrBookAlreadyBorrowed
	}

	book.IsBorrowed = true

	b.data[bookID] = book

	return nil
}

func (b *BooksRepositoryStruct) ReturnBook(input domain.BorrowBookInput) error {
	// Return a borrowed book to the repository

	bookID := input.BookID

	book, ok := b.data[bookID]
	if !ok {
		return domain.ErrBookNotFound
	}

	if !book.IsBorrowed {
		return domain.ErrBookNotBorrowed
	}

	book.IsBorrowed = false

	b.data[bookID] = book

	return nil
}

func (b *BooksRepositoryStruct) PutBook(input domain.PutBookInput) (domain.Book, error) {
	// Update the record

	book, ok := b.data[input.BookID]
	if !ok {
		return domain.Book{}, domain.ErrBookNotFound
	}

	book.Title = input.Title
	book.Author = input.Author

	b.data[input.BookID] = book
	return book, nil
}

func (b *BooksRepositoryStruct) DeleteBook(input domain.DeleteBookInput) error {
	// Delete a book from repository

	bookID := input.BookID

	book, ok := b.data[bookID]
	if !ok {
		return domain.ErrBookNotFound
	}

	if book.IsBorrowed {
		return domain.ErrBookAlreadyBorrowed
	}

	delete(b.data, bookID)

	return nil
}

func (b *BooksRepositoryStruct) ListBooks() []domain.Book {
	// List all books in the repository

	var booksSlc = make([]domain.Book, 0)

	for _, book := range b.data {
		booksSlc = append(booksSlc, book)
	}

	return booksSlc
}
