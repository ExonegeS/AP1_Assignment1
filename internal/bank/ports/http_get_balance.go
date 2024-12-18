package ports

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ExonegeS/AP1_Assignment1/internal/bank/service"
	"github.com/gorilla/mux"
)

func NewEndpointGetBalance(s service.Service, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			walletID = mux.Vars(r)["wallet_id"]
		)

		var input = service.GetBalanceInput{
			WalletID: walletID,
		}

		response, err := s.GetBalance(input)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get balance: %v", err), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		logger.Error("Book recieved successfully", "input", input)
	}
}
