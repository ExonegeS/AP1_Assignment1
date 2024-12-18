package service

import (
	"github.com/ExonegeS/AP1_Assignment1/internal/library/domain"
	"github.com/google/uuid"
)

func (s *service) DeleteBook(input DeleteBookInput) (DeleteBookResponse, error) {
	if err := input.validate(); err != nil {
		return DeleteBookResponse{}, err
	}

	err := s.books.DeleteBook(domain.DeleteBookInput{
		BookID: input.BookID,
	})
	if err != nil {
		return DeleteBookResponse{}, err
	}

	return DeleteBookResponse{Status: domain.BookStatusDeleted}, nil
}

func (q *DeleteBookInput) validate() error {
	if _, err := uuid.Parse(q.BookID); err != nil {
		return err
	}

	return nil
}
