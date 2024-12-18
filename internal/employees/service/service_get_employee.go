package service

import (
	"github.com/ExonegeS/AP1_Assignment1/internal/employees/domain"
)

func (s *service) GetEmployee(input GetEmployeeInput) (GetEmployeeResponse, error) {
	if err := input.validate(); err != nil {
		return GetEmployeeResponse{}, err
	}

	employee, err := s.employees.GetEmployee(domain.GetEmployeeInput{
		EmployeeID: input.EmployeeID,
	})
	if err != nil {
		return GetEmployeeResponse{}, err
	}

	return GetEmployeeResponse{Employee: employee}, nil
}

func (q *GetEmployeeInput) validate() error {
	if q.EmployeeID < 1 {
		return domain.ErrEmployeeInvalidID
	}

	return nil
}
