package service

func (s *service) ListEmployees() ListEmployeesResponse {
	employees := s.employees.ListEmployees()
	return ListEmployeesResponse{Employees: employees}
}
