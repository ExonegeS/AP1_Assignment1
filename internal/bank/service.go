package bank

import (
	"github.com/ExonegeS/AP1_Assignment1/internal/bank/adapters"
	"github.com/ExonegeS/AP1_Assignment1/internal/bank/service"
)

func NewService() service.Service {
	bank := adapters.NewBankRepositoryStruct()
	transaction := adapters.NewTransactionsRepositoryStruct()
	svs := service.NewService(
		bank,
		transaction,
	)

	return svs
}
