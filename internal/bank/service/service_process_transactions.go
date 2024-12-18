package service

import (
	"github.com/ExonegeS/AP1_Assignment1/internal/bank/domain"
)

func (s *service) ProcessTransactions(input ProcessTransactionInput) (ProcessTransactionResponse, error) {
	if err := input.validate(); err != nil {
		return ProcessTransactionResponse{}, err
	}

	transactions := s.transaction.GetTransactions(domain.GetTransactionsInput{
		WalletID: input.WalletID,
		Status:   domain.TransactionStatusWaiting,
		Limit:    0,
	})

	for _, transaction := range transactions {

		if transaction.TransactionStatus != domain.TransactionStatusWaiting {
			continue
		}

		switch true {
		case transaction.Amount > 0:
			if _, err := s.bank.UpdateBalance(domain.UpdateBalanceInput{
				WalletID: input.WalletID,
				Amount:   transaction.Amount,
			}); err != nil {
				continue
			}

			s.transaction.Deposit(domain.DepositTransactionInput{
				TransactionID: transaction.ID,
				WalletID:      input.WalletID,
				Status:        domain.TransactionStatusCompleted,
			})
		case transaction.Amount < 0:
			if _, err := s.bank.UpdateBalance(domain.UpdateBalanceInput{
				WalletID: input.WalletID,
				Amount:   transaction.Amount,
			}); err != nil {
				continue
			}
			s.transaction.Withdraw(domain.WithdrawTransactionInput{
				TransactionID: transaction.ID,
				WalletID:      input.WalletID,
				Status:        domain.TransactionStatusCompleted,
			})
		default:
			continue
		}
	}

	return ProcessTransactionResponse{Status: domain.TransactionStatusCompleted}, nil
}

func (q *ProcessTransactionInput) validate() error {

	return nil
}
