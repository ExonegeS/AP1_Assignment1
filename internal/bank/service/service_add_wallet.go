package service

import (
	"github.com/ExonegeS/AP1_Assignment1/internal/bank/domain"
	"github.com/google/uuid"
)

func (s *service) AddWallet(input AddWalletInput) (AddWalletResponse, error) {
	if err := input.validate(); err != nil {
		return AddWalletResponse{}, err
	}

	input.WalletID = uuid.NewString()
	if err := s.bank.AddWallet(domain.AddWalletInput{
		WalletID:      input.WalletID,
		AccountNumber: input.AccountNumber,
		HolderName:    input.HolderName,
	}); err != nil {
		return AddWalletResponse{}, err
	}

	wallet, err := s.bank.GetWallet(domain.GetWalletInput{
		WalletID: input.WalletID,
	})
	if err != nil {
		return AddWalletResponse{}, err
	}

	return AddWalletResponse{Wallet: wallet}, nil
}

func (q *AddWalletInput) validate() error {
	if !ValidateAccountNumber(q.AccountNumber) {
		return domain.ErrWalletInvalidAccountNumber
	}
	if q.HolderName == "" {
		return domain.ErrWalletInvalidHolderName
	}

	return nil
}

// implement Lung algorithm to check creditcard validation
// Validate checks if the account number is valid using the Luhn algorithm
func ValidateAccountNumber(num string) bool {
	sum := 0
	alternate := false

	for i := len(num) - 1; i >= 0; i-- {
		digit := int(num[i] - '0')
		if digit < 0 || digit > 9 {
			return false // Invalid character
		}
		if alternate {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		alternate = !alternate
	}

	return sum%10 == 0
}
