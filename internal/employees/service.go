package employees

import (
	"github.com/ExonegeS/AP1_Assignment1/internal/employees/adapters"
	"github.com/ExonegeS/AP1_Assignment1/internal/employees/service"
)

func NewService() service.Service {
	employees := adapters.NewEmployeesRepositoryStruct()
	svs := service.NewService(
		employees,
	)

	return svs
}
