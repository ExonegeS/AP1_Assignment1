package service

import (
	"github.com/ExonegeS/AP1_Assignment1/internal/bank/domain"
	"github.com/google/uuid"
)

func (s *service) GetTransactions(input GetTransactionsInput) (GetTransactionsResponse, error) {
	if err := input.validate(); err != nil {
		return GetTransactionsResponse{}, err
	}

	transactions := s.transaction.GetTransactions(domain.GetTransactionsInput{
		WalletID: input.WalletID,
		Status:   domain.TransactionStatusWaiting,
		Limit:    0,
	})

	return GetTransactionsResponse{Transactions: transactions}, nil
}

func (q *GetTransactionsInput) validate() error {
	if _, err := uuid.Parse(q.WalletID); err != nil {
		return err
	}
	return nil
}
