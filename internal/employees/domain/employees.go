package domain

import "fmt"

type EmployeeStatus string

const (
	EmployeeStatusToDo    EmployeeStatus = "TO_DO"
	EmployeeStatusDeleted EmployeeStatus = "DELETED"
)

type EmployeeType string

const (
	EmployeeTypeToDo     EmployeeType = "TO_DO"
	EmployeeTypeFullTime EmployeeType = "FULLTIME"
	EmployeeTypePartTime EmployeeType = "PARTTIME"
)

type Employee struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Salary      uint32  `json:"salary"`
	HourlyRate  uint64  `json:"hourly_rate"`
	HoursWorked float32 `json:"hours_worked"`
}

type EmployeesRepository interface {
	AddEmployee(input AddEmployeeInput) error
	GetEmployee(input GetEmployeeInput) (Employee, error)
	GetEmployeeDetails(input GetEmployeeInput) (string, error)
	PutEmployee(input PutEmployeeInput) (Employee, error)
	DeleteEmployee(input DeleteEmployeInput) error
	ListEmployees() []Employee
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
	GetEmployeeInput struct {
		EmployeeID uint64
	}
	PutEmployeeInput struct {
		EmployeeID   uint64  `json:"id"`
		Name         string  `json:"name"`
		Salary       uint32  `json:"salary"`
		HourlyRate   uint64  `json:"hourly_rate"`
		HoursWorked  float32 `json:"hours_worked"`
		EmployeeType string  `json:"employee_type"`
	}
	DeleteEmployeInput struct {
		EmployeeID uint64
	}
)

var (
	ErrEmployeeExists             = fmt.Errorf("employee with this id already exists")
	ErrEmployeeCannotCreate       = fmt.Errorf("employee cannot be created")
	ErrEmployeeNotFound           = fmt.Errorf("employee not found")
	ErrEmployeeInvalidID          = fmt.Errorf("employee id is invalid")
	ErrEmployeeNameRequired       = fmt.Errorf("employee name is required")
	ErrEmployeeSalaryInvalid      = fmt.Errorf("employee salary is invalid")
	ErrEmployeeHourlyRateInvalid  = fmt.Errorf("employee hourly rate is invalid")
	ErrEmployeeHoursWorkedInvalid = fmt.Errorf("employee hours worked is invalid")
	ErrEmployeeTypeInvalid        = fmt.Errorf("employee type is invalid")
)
