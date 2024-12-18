package service

import "github.com/ExonegeS/AP1_Assignment1/internal/employees/domain"

type Service interface {
	AddEmployee(input AddEmployeeInput) (AddEmployeeResponse, error)
	GetEmployee(input GetEmployeeInput) (GetEmployeeResponse, error)
	GetEmployeeDetails(input GetEmployeeInput) (GetEmployeeDetailsResponse, error)
	PutEmployee(input PutEmployeeInput) (PutEmployeeResponse, error)
	DeleteEmployee(input DeleteEmployeInput) (DeleteEmployeeResponse, error)
	ListEmployees() ListEmployeesResponse
}

type (
	AddEmployeeInput struct {
		EmployeeID   uint64  `json:"id"`
		Name         string  `json:"name"`
		Salary       uint32  `json:"salary"`
		HourlyRate   uint64  `json:"hourly_rate"`
		HoursWorked  float32 `json:"hours_worked"`
		EmployeeType string  `json:"employee_type"`
	}
	AddEmployeeResponse struct {
		Employee domain.Employee `json:"employee"`
	}
)

type (
	GetEmployeeInput struct {
		EmployeeID uint64
	}
	GetEmployeeResponse struct {
		Employee domain.Employee `json:"employee"`
	}
	GetEmployeeDetailsResponse struct {
		Details string `json:"details"`
	}
)

type (
	PutEmployeeInput struct {
		EmployeeID   uint64  `json:"id"`
		Name         string  `json:"name"`
		Salary       uint32  `json:"salary"`
		HourlyRate   uint64  `json:"hourly_rate"`
		HoursWorked  float32 `json:"hours_worked"`
		EmployeeType string  `json:"employee_type"`
	}
	PutEmployeeResponse struct {
		Employee domain.Employee `json:"employee"`
	}
)

type (
	DeleteEmployeInput struct {
		EmployeeID uint64
	}
	DeleteEmployeeResponse struct {
		Status domain.EmployeeStatus
	}
)

type (
	ListEmployeesResponse struct {
		Employees []domain.Employee `json:"employees"`
	}
)

type service struct {
	employees domain.EmployeesRepository
}

var _ Service = (*service)(nil)

func NewService(employees domain.EmployeesRepository) Service {
	return &service{
		employees: employees,
	}
}
