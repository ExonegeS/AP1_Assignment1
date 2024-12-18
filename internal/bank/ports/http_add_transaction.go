package ports

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ExonegeS/AP1_Assignment1/internal/bank/service"
	"github.com/gorilla/mux"
)

func NewEndpointAddTransaction(s service.Service, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			walletID = mux.Vars(r)["wallet_id"]
		)
		var input service.AddTransactionInput

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			logger.Error("Invalid input", "input", input)
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		if _, err := s.GetWallet(service.GetWalletInput{WalletID: walletID}); err != nil {
			logger.Error("Wallet not found", "wallet_id", walletID)
			http.Error(w, fmt.Sprintf("Failed to add wallet: %v", err), http.StatusInternalServerError)
			return
		}

		input.WalletID = walletID
		transaction, err := s.AddTransaction(input)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to add transaction: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(transaction)

		logger.Error("Employee added successfully", "input", input)
	}
}
