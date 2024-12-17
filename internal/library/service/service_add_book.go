package service

import (
	"github.com/ExonegeS/AP1_Assignment1/internal/library/domain"
	"github.com/google/uuid"
)

func (s *service) AddBook(input AddBookInput) (AddBookResponse, error) {
	if err := input.validate(); err != nil {
		return AddBookResponse{}, err
	}

	bookID := uuid.NewString()
	if err := s.books.AddBook(domain.AddBookInput{
		BookID: bookID,
		Title:  input.Title,
		Author: input.Author,
	}); err != nil {
		return AddBookResponse{}, err
	}

	book, err := s.books.GetBook(domain.GetBookInput{
		BookID: bookID,
	})
	if err != nil {
		return AddBookResponse{}, err
	}

	return AddBookResponse{Book: book}, nil
}

func (q *AddBookInput) validate() error {
	if q.Title == "" {
		return domain.ErrBookInvalidInputTitle
	}

	if q.Author == "" {
		return domain.ErrBookInvalidInputAuthor
	}

	return nil
}
