package domain

import (
	"fmt"
	"time"
)

type TransactionStatus string

const (
	TransactionStatusToDo      TransactionStatus = "TO_DO"
	TransactionStatusAny       TransactionStatus = "ANY"
	TransactionStatusWaiting   TransactionStatus = "WAITING"
	TransactionStatusCompleted TransactionStatus = "COMPLETED"
	TransactionStatusCanceled  TransactionStatus = "CANCELED"
)

type TransactionType string

const (
	TransactionTypeToDo TransactionType = "TO_DO"
)

type Transaction struct {
	ID                string            `json:"id"`
	WalletID          string            `json:"wallet_id"`
	TransactionStatus TransactionStatus `json:"status"`
	Amount            float64           `json:"amount"`
	CreationTime      time.Time
}

type TransactionRepository interface {
	AddTransaction(input AddTransactionInput) error
	GetTransactions(input GetTransactionsInput) []Transaction
	Deposit(input DepositTransactionInput) (TransactionStatus, error)
	Withdraw(input WithdrawTransactionInput) (TransactionStatus, error)
}

type (
	AddTransactionInput struct {
		TransactionID string  `json:"id"`
		WalletID      string  `json:"wallet_id"`
		Amount        float64 `json:"amount"`
		CreationTime  time.Time
	}
	ProcessTransactionInput struct {
		Wallet       Wallet        `json:"wallet"`
		Transactions []Transaction `json:"transactions"`
	}
	GetTransactionsInput struct {
		WalletID string            `json:"wallet_id"`
		Status   TransactionStatus `json:"status"`
		Limit    uint              `json:"limit"`
	}
	DepositTransactionInput struct {
		TransactionID string            `json:"id"`
		WalletID      string            `json:"wallet_id"`
		Status        TransactionStatus `json:"status"`
	}
	WithdrawTransactionInput struct {
		TransactionID string            `json:"id"`
		WalletID      string            `json:"wallet_id"`
		Status        TransactionStatus `json:"status"`
	}
)

var (
	ErrTransactionExists        = fmt.Errorf("transaction with this id already exists")
	ErrTransactionCannotCreate  = fmt.Errorf("transaction cannot be created")
	ErrTransactionCannotProcess = fmt.Errorf("transaction cannot be processed")
	ErrTransactionNotFound      = fmt.Errorf("transaction not found")
	ErrTransactionInvalidID     = fmt.Errorf("transaction id is invalid")
	ErrTransactionInvalidAmount = fmt.Errorf("amount must be non-zero")
)
