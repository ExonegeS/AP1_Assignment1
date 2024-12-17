package ports

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/ExonegeS/AP1_Assignment1/internal/library/service"
)

func NewEndpointGetBooks(s service.Service, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := s.ListBooks()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		logger.Error("Books recieved successfully")
	}
}
