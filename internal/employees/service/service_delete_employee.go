package service

import (
	"github.com/ExonegeS/AP1_Assignment1/internal/employees/domain"
)

func (s *service) DeleteEmployee(input DeleteEmployeInput) (DeleteEmployeeResponse, error) {
	if err := input.validate(); err != nil {
		return DeleteEmployeeResponse{}, err
	}

	err := s.employees.DeleteEmployee(domain.DeleteEmployeInput{
		EmployeeID: input.EmployeeID,
	})
	if err != nil {
		return DeleteEmployeeResponse{}, err
	}

	return DeleteEmployeeResponse{Status: domain.EmployeeStatusDeleted}, nil
}

func (q *DeleteEmployeInput) validate() error {
	if q.EmployeeID < 1 {
		return domain.ErrEmployeeInvalidID
	}

	return nil
}
