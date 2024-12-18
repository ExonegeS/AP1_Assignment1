package service

import "github.com/ExonegeS/AP1_Assignment1/internal/bank/domain"

type Service interface {
	AddWallet(input AddWalletInput) (AddWalletResponse, error)
	GetBalance(input GetBalanceInput) (GetBalanceResponse, error)
	GetWallet(input GetWalletInput) (GetWalletResponse, error)

	AddTransaction(input AddTransactionInput) (AddTransactionResponse, error)
	GetTransactions(input GetTransactionsInput) (GetTransactionsResponse, error)
	ProcessTransactions(input ProcessTransactionInput) (ProcessTransactionResponse, error)
	// Deposit(input DepositTransactionInput) (ProcessTransactionResponse, error)
	// Withdraw(input WithdrawTransactionInput) (ProcessTransactionResponse, error)
}

type (
	AddWalletInput struct {
		WalletID      string `json:"wallet_id"`
		AccountNumber string `json:"account_number"`
		HolderName    string `json:"holder_name"`
	}
	AddWalletResponse struct {
		Wallet domain.Wallet `json:"wallet"`
	}
)

type (
	GetWalletInput struct {
		WalletID string
	}
	GetWalletResponse struct {
		Wallet domain.Wallet `json:"wallet"`
	}
)

type (
	GetBalanceInput struct {
		WalletID string
	}
	GetBalanceResponse struct {
		Balance domain.Balance `json:"balance"`
	}
)
type (
	GetTransactionsInput struct {
		WalletID string                   `json:"wallet_id"`
		Status   domain.TransactionStatus `json:"status"`
		Limit    uint                     `json:"limit"`
	}
	GetTransactionsResponse struct {
		Transactions []domain.Transaction `json:"transactions"`
	}
)
type (
	AddTransactionInput struct {
		TransactionID string  `json:"id"`
		WalletID      string  `json:"wallet_id"`
		Amount        float64 `json:"amount"`
	}
	AddTransactionResponse struct {
		TotalTransaction int `json:"transaction_length"`
	}
)
type (
	ProcessTransactionInput struct {
		TransactionID string  `json:"id"`
		WalletID      string  `json:"wallet_id"`
		Amount        float64 `json:"amount"`
	}
	ProcessTransactionResponse struct {
		Status domain.TransactionStatus `json:"status"`
	}
)

type (
	DepositTransactionInput struct {
		TransactionID string                   `json:"id"`
		Status        domain.TransactionStatus `json:"status"`
		Amount        float64                  `json:"amount"`
	}
	WithdrawTransactionInput struct {
		TransactionID string                   `json:"id"`
		Status        domain.TransactionStatus `json:"status"`
	}
)

type service struct {
	bank        domain.BankRepository
	transaction domain.TransactionRepository
}

var _ Service = (*service)(nil)

func NewService(bank domain.BankRepository, transaction domain.TransactionRepository) Service {
	return &service{
		bank:        bank,
		transaction: transaction,
	}
}
