package ports

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/ExonegeS/AP1_Assignment1/internal/employees/service"
	"github.com/gorilla/mux"
)

func NewEndpointPutEmployee(s service.Service, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			employeeIDStr = mux.Vars(r)["employee_id"]
		)

		employeeID, err := strconv.ParseUint(employeeIDStr, 10, 64)
		var input = service.PutEmployeeInput{
			EmployeeID: employeeID,
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			logger.Error("Invalid input", "input", input)
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		if err != nil {
			logger.Error("Failed to parse employee ID", "err", err)
			http.Error(w, "Invalid employee ID", http.StatusBadRequest)
			return
		}

		response, err := s.PutEmployee(input)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to update employee: %v", err), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		logger.Error("Employee recieved successfully", "input", input)
	}
}
