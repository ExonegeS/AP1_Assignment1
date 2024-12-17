package service

import (
	"github.com/ExonegeS/AP1_Assignment1/internal/library/domain"
	"github.com/google/uuid"
)

func (s *service) GetBook(input GetBookInput) (GetBookResponse, error) {
	if err := input.validate(); err != nil {
		return GetBookResponse{}, err
	}

	book, err := s.books.GetBook(domain.GetBookInput{
		BookID: input.BookID,
	})
	if err != nil {
		return GetBookResponse{}, err
	}

	return GetBookResponse{Book: book}, nil
}

func (q *GetBookInput) validate() error {
	if _, err := uuid.Parse(q.BookID); err != nil {
		return err
	}

	return nil
}
