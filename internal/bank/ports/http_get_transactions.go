package ports

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ExonegeS/AP1_Assignment1/internal/bank/domain"
	"github.com/ExonegeS/AP1_Assignment1/internal/bank/service"
	"github.com/gorilla/mux"
)

func NewEndpointGetTransactions(s service.Service, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			walletID = mux.Vars(r)["wallet_id"]
		)

		var input = service.GetTransactionsInput{
			WalletID: walletID,
			Status:   domain.TransactionStatusAny,
		}

		response, err := s.GetTransactions(input)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get transactions: %v", err), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		logger.Error("Book recieved successfully", "input", input)
	}
}
