package service

import (
	"github.com/ExonegeS/AP1_Assignment1/internal/employees/domain"
)

func (s *service) AddEmployee(input AddEmployeeInput) (AddEmployeeResponse, error) {
	if err := input.validate(); err != nil {
		return AddEmployeeResponse{}, err
	}

	if err := s.employees.AddEmployee(domain.AddEmployeeInput{
		EmployeeID:   input.EmployeeID,
		Name:         input.Name,
		Salary:       input.Salary,
		HourlyRate:   input.HourlyRate,
		HoursWorked:  input.HoursWorked,
		EmployeeType: input.EmployeeType,
	}); err != nil {
		return AddEmployeeResponse{}, err
	}

	employee, err := s.employees.GetEmployee(domain.GetEmployeeInput{
		EmployeeID: input.EmployeeID,
	})
	if err != nil {
		return AddEmployeeResponse{}, err
	}

	return AddEmployeeResponse{Employee: employee}, nil
}

func (q *AddEmployeeInput) validate() error {
	if q.EmployeeID < 1 { // no value of type uint64 is less than 0 (SA4003
		return domain.ErrEmployeeInvalidID
	}

	if q.Name == "" {
		return domain.ErrEmployeeNameRequired
	}

	switch q.EmployeeType {
	case string(domain.EmployeeTypeFullTime):
		if q.Salary < 1 { // no value of type uint32 is less than 0 (SA4003
			return domain.ErrEmployeeSalaryInvalid
		}
	case string(domain.EmployeeTypePartTime):
		if q.HourlyRate < 1 {
			return domain.ErrEmployeeHourlyRateInvalid
		}
		if q.HoursWorked < 0 {
			return domain.ErrEmployeeHoursWorkedInvalid
		}
	default:
		return domain.ErrEmployeeTypeInvalid
	}

	return nil
}
