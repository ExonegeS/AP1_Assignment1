package service

import (
	"github.com/ExonegeS/AP1_Assignment1/internal/bank/domain"
)

func (s *service) GetBalance(input GetBalanceInput) (GetBalanceResponse, error) {
	if err := input.validate(); err != nil {
		return GetBalanceResponse{}, err
	}

	balance, err := s.bank.GetBalance(domain.GetBalanceInput{
		WalletID: input.WalletID,
	})
	if err != nil {
		return GetBalanceResponse{}, err
	}

	return GetBalanceResponse{Balance: balance}, nil
}

func (q *GetBalanceInput) validate() error {
	return nil
}
