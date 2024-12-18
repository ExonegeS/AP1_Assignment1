package service

import (
	"time"

	"github.com/ExonegeS/AP1_Assignment1/internal/bank/domain"
	"github.com/google/uuid"
)

func (s *service) AddTransaction(input AddTransactionInput) (AddTransactionResponse, error) {
	if err := input.validate(); err != nil {
		return AddTransactionResponse{}, err
	}
	transactionID := uuid.NewString()

	if err := s.transaction.AddTransaction(domain.AddTransactionInput{
		TransactionID: transactionID,
		WalletID:      input.WalletID,
		Amount:        input.Amount,
		CreationTime:  time.Now(),
	}); err != nil {
		return AddTransactionResponse{}, err
	}

	transactions := s.transaction.GetTransactions(domain.GetTransactionsInput{
		WalletID: input.WalletID,
		Status:   domain.TransactionStatusWaiting,
		Limit:    0,
	})

	return AddTransactionResponse{TotalTransaction: len(transactions)}, nil
}

func (q *AddTransactionInput) validate() error {
	if q.Amount == 0 {
		return domain.ErrTransactionInvalidAmount
	}

	if _, err := uuid.Parse(q.WalletID); err != nil {
		return err
	}
	
	return nil
}
