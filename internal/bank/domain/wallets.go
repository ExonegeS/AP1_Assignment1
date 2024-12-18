package domain

import "fmt"

type WalletStatus string

const (
	WalletStatusToDo           WalletStatus = "TO_DO"
	WalletStatusError          WalletStatus = "SOME_ERROR"
	WalletStatusBalanceUpdated WalletStatus = "BALANCE_UPDATED"
)

type WalletType string

const (
	WalletTypeToDo     WalletType = "TO_DO"
	WalletTypeDebit    WalletType = "DEBIT"
	WalletTypeToCredit WalletType = "CREDIT"
)

type Wallet struct {
	ID            string  `json:"id"`
	AccountNumber string  `json:"account_number"`
	HolderName    string  `json:"holder_name"`
	Balance       Balance `json:"balance"`
}

type Balance struct {
	Amount float64 `json:"amount"`
}

type BankRepository interface {
	AddWallet(input AddWalletInput) error
	GetWallet(input GetWalletInput) (Wallet, error)
	GetBalance(input GetBalanceInput) (Balance, error)
	UpdateBalance(input UpdateBalanceInput) (Balance, error)
}

type (
	AddWalletInput struct {
		WalletID      string `json:"wallet_id"`
		AccountNumber string `json:"account_number"`
		HolderName    string `json:"holder_name"`
	}
	GetWalletInput struct {
		WalletID string
	}
	GetBalanceInput struct {
		WalletID string
	}
	UpdateBalanceInput struct {
		WalletID string  `json:"wallet_id"`
		Amount   float64 `json:"amount"`
	}
)

var (
	ErrWalletExists               = fmt.Errorf("wallet with this id already exists")
	ErrWalletCannotCreate         = fmt.Errorf("wallet cannot be created")
	ErrWalletNotFound             = fmt.Errorf("wallet not found")
	ErrWalletInvalidID            = fmt.Errorf("wallet id is invalid")
	ErrWalletInvalidAccountNumber = fmt.Errorf("account number id is invalid")
	ErrWalletInvalidHolderName    = fmt.Errorf("holder name is invalid")
	ErrInsufficientBalance        = fmt.Errorf("insufficient balance in the wallet")
)
