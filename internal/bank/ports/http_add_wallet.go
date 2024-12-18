package ports

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ExonegeS/AP1_Assignment1/internal/bank/service"
)

func NewEndpointAddWallet(s service.Service, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input service.AddWalletInput

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			logger.Error("Invalid input", "input", input)
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		wallet, err := s.AddWallet(input)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to add wallet: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(wallet)

		logger.Error("Wallet added successfully", "input", input)
	}
}
