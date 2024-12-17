package service

import (
	"github.com/ExonegeS/AP1_Assignment1/internal/library/domain"
	"github.com/google/uuid"
)

func (s *service) BorrowBook(input BorrowBookInput) (BorrowBookResponse, error) {
	if err := input.validate(); err != nil {
		return BorrowBookResponse{}, err
	}

	if err := s.books.BorrowBook(domain.BorrowBookInput{
		BookID: input.BookID,
	}); err != nil {
		return BorrowBookResponse{}, err
	}

	return BorrowBookResponse{Status: domain.BookStatusBooked}, nil
}

func (q *BorrowBookInput) validate() error {
	if _, err := uuid.Parse(q.BookID); err != nil {
		return err
	}

	return nil
}
