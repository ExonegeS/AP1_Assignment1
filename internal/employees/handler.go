package employees

import (
	"log/slog"
	"net/http"

	"github.com/ExonegeS/AP1_Assignment1/internal/employees/ports"
	"github.com/ExonegeS/AP1_Assignment1/internal/employees/service"
	"github.com/gorilla/mux"
)

func NewHandler(s service.Service, logger *slog.Logger) http.Handler {
	r := mux.NewRouter()
	r.Handle("/employees/add", ports.NewEndpointAddEmployee(s, logger)).Methods(http.MethodPost)
	r.Handle("/employees/list", ports.NewEndpointGetEmployees(s, logger)).Methods(http.MethodGet)
	r.Handle("/employees/{employee_id}/details", ports.NewEndpointGetEmployeeDetails(s, logger)).Methods(http.MethodGet)
	r.Handle("/employees/{employee_id}", ports.NewEndpointGetEmployee(s, logger)).Methods(http.MethodGet)
	r.Handle("/employees/{employee_id}", ports.NewEndpointPutEmployee(s, logger)).Methods(http.MethodPut)
	r.Handle("/employees/{employee_id}", ports.NewEndpointDeleteEmployee(s, logger)).Methods(http.MethodDelete)

	// Register other routes as needed
	return r
}
