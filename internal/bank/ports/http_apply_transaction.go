package ports

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ExonegeS/AP1_Assignment1/internal/bank/service"
	"github.com/gorilla/mux"
)

func NewEndpointApplyTransaction(s service.Service, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			walletID = mux.Vars(r)["wallet_id"]
		)
		var input service.ProcessTransactionInput

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			logger.Error("Invalid input", "input", input)
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		input.WalletID = walletID

		transaction, err := s.ProcessTransactions(input)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to process transactions: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(transaction)

		logger.Error("Employee added successfully", "input", input)
	}
}
