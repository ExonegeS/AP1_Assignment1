package service

import (
	"github.com/ExonegeS/AP1_Assignment1/internal/employees/domain"
)

func (s *service) PutEmployee(input PutEmployeeInput) (PutEmployeeResponse, error) {
	if err := input.validate(); err != nil {
		return PutEmployeeResponse{}, err
	}

	employee, err := s.employees.PutEmployee(domain.PutEmployeeInput{
		EmployeeID:   input.EmployeeID,
		Name:         input.Name,
		Salary:       input.Salary,
		HourlyRate:   input.HourlyRate,
		HoursWorked:  input.HoursWorked,
		EmployeeType: input.EmployeeType,
	})
	if err != nil {
		return PutEmployeeResponse{}, err
	}

	return PutEmployeeResponse{Employee: employee}, nil
}

func (q *PutEmployeeInput) validate() error {
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
