package adapters

import (
	"github.com/ExonegeS/AP1_Assignment1/internal/bank/domain"
)

type BankRepositoryStruct struct {
	data map[string]domain.Wallet
}

var _ domain.BankRepository = (*BankRepositoryStruct)(nil)

func NewBankRepositoryStruct() *BankRepositoryStruct {
	return &BankRepositoryStruct{
		data: make(map[string]domain.Wallet),
	}
}

func (e *BankRepositoryStruct) AddWallet(input domain.AddWalletInput) error {
	// Add a new wallet to the repository

	var newWallet = domain.Wallet{
		ID:            input.WalletID,
		AccountNumber: input.AccountNumber,
		HolderName:    input.HolderName,
		Balance:       domain.Balance{Amount: 0},
	}

	if _, ok := e.data[input.WalletID]; ok {
		return domain.ErrWalletExists
	}

	e.data[input.WalletID] = newWallet

	return nil
}

func (e *BankRepositoryStruct) GetWallet(input domain.GetWalletInput) (domain.Wallet, error) {
	walletID := input.WalletID

	wallet, ok := e.data[walletID]
	if !ok {
		return domain.Wallet{}, domain.ErrWalletNotFound
	}

	return wallet, nil
}
func (e *BankRepositoryStruct) GetBalance(input domain.GetBalanceInput) (domain.Balance, error) {
	// Get a balance from the repository

	walletID := input.WalletID

	wallet, ok := e.data[walletID]
	if !ok {
		return domain.Balance{}, domain.ErrWalletNotFound
	}

	return domain.Balance{Amount: wallet.Balance.Amount}, nil
}

func (e *BankRepositoryStruct) UpdateBalance(input domain.UpdateBalanceInput) (domain.Balance, error) {
	// Update a balance

	wallet, ok := e.data[input.WalletID]
	if !ok {
		return domain.Balance{}, domain.ErrWalletNotFound
	}

	newBalance := wallet.Balance
	if newBalance.Amount + input.Amount < 0 {
		return domain.Balance{}, domain.ErrInsufficientBalance
	}
	newBalance.Amount += input.Amount

	wallet.Balance = newBalance

	e.data[input.WalletID] = wallet

	return wallet.Balance, nil

}
