package service

import (
	"github.com/ExonegeS/AP1_Assignment1/internal/bank/domain"
)

func (s *service) GetWallet(input GetWalletInput) (GetWalletResponse, error) {
	if err := input.validate(); err != nil {
		return GetWalletResponse{}, err
	}

	wallet, err := s.bank.GetWallet(domain.GetWalletInput{
		WalletID: input.WalletID,
	})
	if err != nil {
		return GetWalletResponse{}, err
	}

	return GetWalletResponse{Wallet: wallet}, nil
}

func (q *GetWalletInput) validate() error {
	return nil
}
