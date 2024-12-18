package adapters

import (
	"github.com/ExonegeS/AP1_Assignment1/internal/employees/domain"
)

type EmployeesRepositoryStruct struct {
	data map[uint64]Employee
}

var _ domain.EmployeesRepository = (*EmployeesRepositoryStruct)(nil)

func NewEmployeesRepositoryStruct() *EmployeesRepositoryStruct {
	return &EmployeesRepositoryStruct{
		data: make(map[uint64]Employee),
	}
}

func (e *EmployeesRepositoryStruct) AddEmployee(input domain.AddEmployeeInput) error {
	// Add a new employee to the repository

	var newEmployee Employee

	switch input.EmployeeType {
	case string(domain.EmployeeTypeFullTime):
		newEmployee = &FullTimeEmployee{
			ID:     input.EmployeeID,
			Name:   input.Name,
			Salary: input.Salary,
		}
	case string(domain.EmployeeTypePartTime):
		newEmployee = &PartTimeEmployee{
			ID:          input.EmployeeID,
			Name:        input.Name,
			HourlyRate:  input.HourlyRate,
			HoursWorked: input.HoursWorked,
		}
	default:
		return domain.ErrEmployeeTypeInvalid
	}

	if _, ok := e.data[input.EmployeeID]; ok {
		return domain.ErrEmployeeExists
	}

	e.data[input.EmployeeID] = newEmployee
	return nil
}

func (e *EmployeesRepositoryStruct) GetEmployee(input domain.GetEmployeeInput) (domain.Employee, error) {
	// Get a employee from the repository

	employeeID := input.EmployeeID

	employee, ok := e.data[employeeID]
	if !ok {
		return domain.Employee{}, domain.ErrEmployeeNotFound
	}

	return employee.ConvertToDomain(), nil
}

func (e *EmployeesRepositoryStruct) GetEmployeeDetails(input domain.GetEmployeeInput) (string, error) {
	// Get a employee readable details

	employeeID := input.EmployeeID

	employee, ok := e.data[employeeID]
	if !ok {
		return "", domain.ErrEmployeeNotFound
	}

	return employee.GetDetails(), nil
}

func (e *EmployeesRepositoryStruct) PutEmployee(input domain.PutEmployeeInput) (domain.Employee, error) {
	// Update the record
	var employee Employee

	_, ok := e.data[input.EmployeeID]
	if !ok {
		return domain.Employee{}, domain.ErrEmployeeNotFound
	}

	switch input.EmployeeType {
	case string(domain.EmployeeTypeFullTime):
		employee = &FullTimeEmployee{
			ID:     input.EmployeeID,
			Name:   input.Name,
			Salary: input.Salary,
		}
	case string(domain.EmployeeTypePartTime):
		employee = &PartTimeEmployee{
			ID:          input.EmployeeID,
			Name:        input.Name,
			HourlyRate:  input.HourlyRate,
			HoursWorked: input.HoursWorked,
		}
	default:
		return domain.Employee{}, domain.ErrEmployeeTypeInvalid
	}

	e.data[input.EmployeeID] = employee
	return employee.ConvertToDomain(), nil
}

func (e *EmployeesRepositoryStruct) DeleteEmployee(input domain.DeleteEmployeInput) error {
	// Delete a employee from repository

	employeeID := input.EmployeeID

	_, ok := e.data[employeeID]
	if !ok {
		return domain.ErrEmployeeNotFound
	}

	delete(e.data, employeeID)

	return nil
}

func (e *EmployeesRepositoryStruct) ListEmployees() []domain.Employee {
	// List all employees in the repository

	var employeesSlc = make([]domain.Employee, 0)

	for _, employee := range e.data {
		employeesSlc = append(employeesSlc, employee.ConvertToDomain())
	}

	return employeesSlc
}
