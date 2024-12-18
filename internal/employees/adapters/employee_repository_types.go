package adapters

import (
	"fmt"

	"github.com/ExonegeS/AP1_Assignment1/internal/employees/domain"
)

type Employee interface {
	GetDetails() string
	ConvertToDomain() domain.Employee
}

type FullTimeEmployee struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	Salary uint32 `json:"salary"`
	Employee
}

type PartTimeEmployee struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	HourlyRate  uint64  `json:"hourly_rate"`
	HoursWorked float32 `json:"hours_worked"`
	Employee
}

func (f *FullTimeEmployee) GetDetails() string {
	var details string = fmt.Sprintf("FullTime Worker #%v\nFUll Name: '%s'\nSalary: %v KZT",
		f.ID,
		f.Name,
		f.Salary,
	)
	return details
}

func (p *PartTimeEmployee) GetDetails() string {
	var details string = fmt.Sprintf("PartTime Worker #%v\nFUll Name: %s\nSalary/Hour: %v KZT | Worked: %v h | Earned: %v KZT",
		p.ID,
		p.Name,
		p.HourlyRate,
		p.HoursWorked,
		float32(p.HourlyRate)*p.HoursWorked,
	)
	return details
}

func (f *FullTimeEmployee) ConvertToDomain() domain.Employee {
	return domain.Employee{
		ID:     f.ID,
		Name:   f.Name,
		Salary: f.Salary,
	}
}

func (p *PartTimeEmployee) ConvertToDomain() domain.Employee {
	return domain.Employee{
		ID:          p.ID,
		Name:        p.Name,
		HourlyRate:  p.HourlyRate,
		HoursWorked: p.HoursWorked,
	}
}
