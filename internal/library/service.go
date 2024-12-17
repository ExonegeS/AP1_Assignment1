package library

import (
	"github.com/ExonegeS/AP1_Assignment1/internal/library/adapters"
	"github.com/ExonegeS/AP1_Assignment1/internal/library/service"
)

func NewService() service.Service {
	books := adapters.NewBooksRepositoryStruct()
	svs := service.NewService(
		books,
	)

	return svs
}
