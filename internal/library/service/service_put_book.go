package service

import (
	"github.com/ExonegeS/AP1_Assignment1/internal/library/domain"
	"github.com/google/uuid"
)

func (s *service) PutBook(input PutBookInput) (PutBookResponse, error) {
	if err := input.validate(); err != nil {
		return PutBookResponse{}, err
	}

	book, err := s.books.PutBook(domain.PutBookInput{
		BookID: input.BookID,
		Title:  input.Title,
		Author: input.Author,
	})
	if err != nil {
		return PutBookResponse{}, err
	}

	return PutBookResponse{Book: book}, nil
}

func (q *PutBookInput) validate() error {
	if _, err := uuid.Parse(q.BookID); err != nil {
		return err
	}

	if q.Title == "" {
		return domain.ErrBookInvalidInputTitle
	}

	if q.Author == "" {
		return domain.ErrBookInvalidInputAuthor
	}

	return nil
}
