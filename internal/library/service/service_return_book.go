package service

import (
	"github.com/ExonegeS/AP1_Assignment1/internal/library/domain"
)

func (s *service) ReturnBook(input BorrowBookInput) (BorrowBookResponse, error) {
	if err := input.validate(); err != nil {
		return BorrowBookResponse{}, err
	}

	if err := s.books.ReturnBook(domain.BorrowBookInput{
		BookID: input.BookID,
	}); err != nil {
		return BorrowBookResponse{}, err
	}

	return BorrowBookResponse{Status: domain.BookStatusAvaiable}, nil
}
