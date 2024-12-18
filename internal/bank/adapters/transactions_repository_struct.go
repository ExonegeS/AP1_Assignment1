package adapters

import (
	"github.com/ExonegeS/AP1_Assignment1/internal/bank/domain"
)

type TransactionRepositoryStuct struct {
	data map[string][]domain.Transaction // map[wallet_id]slice of wallet transactions
}

var _ domain.TransactionRepository = (*TransactionRepositoryStuct)(nil)

func NewTransactionsRepositoryStruct() *TransactionRepositoryStuct {
	return &TransactionRepositoryStuct{
		data: make(map[string][]domain.Transaction),
	}
}

func (e *TransactionRepositoryStuct) AddTransaction(input domain.AddTransactionInput) error {
	// Add a new transaction to the repository

	var newTransaction = domain.Transaction{
		ID:                input.TransactionID,
		WalletID:          input.WalletID,
		TransactionStatus: domain.TransactionStatusWaiting,
		Amount:            input.Amount,
		CreationTime:      input.CreationTime,
	}

	if _, ok := e.data[input.WalletID]; !ok {
		e.data[input.WalletID] = make([]domain.Transaction, 0)
	}

	e.data[input.WalletID] = append(e.data[input.WalletID], newTransaction)

	return nil
}

func (e *TransactionRepositoryStuct) GetTransactions(input domain.GetTransactionsInput) []domain.Transaction {
	transactions, ok := e.data[input.WalletID]
	if !ok {
		return nil
	}

	if input.Status != domain.TransactionStatusAny {
		transactions = filterTransactionsByStatus(transactions, input.Status)
	}
	if input.Limit > 0 {
		return transactions[:input.Limit]
	}

	return transactions
}

func (e *TransactionRepositoryStuct) Deposit(input domain.DepositTransactionInput) (domain.TransactionStatus, error) {
	// Update wallet balance and record the transaction

	transactions, ok := e.data[input.WalletID]
	if !ok {
		return domain.TransactionStatusCanceled, domain.ErrWalletNotFound
	}

	// Find and update the transaction
	for i, transaction := range transactions {
		if transaction.ID == input.TransactionID {
			// Update the status to Completed
			e.data[input.WalletID][i].TransactionStatus = domain.TransactionStatusCompleted

			return domain.TransactionStatusCompleted, nil
		}
	}

	return domain.TransactionStatusCanceled, domain.ErrTransactionNotFound
}

func (e *TransactionRepositoryStuct) Withdraw(input domain.WithdrawTransactionInput) (domain.TransactionStatus, error) {
	// Update wallet balance and record the transaction

	transactions, ok := e.data[input.WalletID]
	if !ok {
		return domain.TransactionStatusCanceled, domain.ErrWalletNotFound
	}

	// Find and update the transaction
	for i, transaction := range transactions {
		if transaction.ID == input.TransactionID {
			// Update the status to Completed
			e.data[input.WalletID][i].TransactionStatus = domain.TransactionStatusCompleted

			return domain.TransactionStatusCompleted, nil
		}
	}

	return domain.TransactionStatusCanceled, domain.ErrTransactionNotFound
}

func filterTransactionsByStatus(transactions []domain.Transaction, status domain.TransactionStatus) []domain.Transaction {
	var filtered []domain.Transaction

	for _, t := range transactions {
		if t.TransactionStatus == status {
			filtered = append(filtered, t)
		}
	}
	return filtered
}
