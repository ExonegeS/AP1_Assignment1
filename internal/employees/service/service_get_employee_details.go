package service

import (
	"github.com/ExonegeS/AP1_Assignment1/internal/employees/domain"
)

func (s *service) GetEmployeeDetails(input GetEmployeeInput) (GetEmployeeDetailsResponse, error) {
	if err := input.validate(); err != nil {
		return GetEmployeeDetailsResponse{}, err
	}

	details, err := s.employees.GetEmployeeDetails(domain.GetEmployeeInput{
		EmployeeID: input.EmployeeID,
	})
	if err != nil {
		return GetEmployeeDetailsResponse{}, err
	}

	return GetEmployeeDetailsResponse{Details: details}, nil
}
