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

func NewEndpointGetEmployee(s service.Service, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			employeeIDStr = mux.Vars(r)["employee_id"]
		)

		employeeID, err := strconv.ParseUint(employeeIDStr, 10, 64)
		if err != nil {
			logger.Error("Failed to parse employee ID", "err", err)
			http.Error(w, "Invalid employee ID", http.StatusBadRequest)
			return
		}

		var input = service.GetEmployeeInput{
			EmployeeID: employeeID,
		}

		response, err := s.GetEmployee(input)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get employee: %v", err), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		logger.Error("Employee recieved successfully", "input", input)
	}
}
