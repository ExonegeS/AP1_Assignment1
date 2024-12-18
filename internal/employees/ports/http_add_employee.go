package ports

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ExonegeS/AP1_Assignment1/internal/employees/service"
)

func NewEndpointAddEmployee(s service.Service, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input service.AddEmployeeInput

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			logger.Error("Invalid input", "input", input)
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		employee, err := s.AddEmployee(input)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to add employee: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employee)

		logger.Error("Employee added successfully", "input", input)
	}
}
