package bank

import (
	"log/slog"
	"net/http"

	"github.com/ExonegeS/AP1_Assignment1/internal/bank/ports"
	"github.com/ExonegeS/AP1_Assignment1/internal/bank/service"
	"github.com/gorilla/mux"
)

func NewHandler(s service.Service, logger *slog.Logger) http.Handler {
	r := mux.NewRouter()
	r.Handle("/bank/add", ports.NewEndpointAddWallet(s, logger)).Methods(http.MethodPost)
	r.Handle("/bank/{wallet_id}/balance", ports.NewEndpointGetBalance(s, logger)).Methods(http.MethodGet)
	r.Handle("/bank/{wallet_id}/add", ports.NewEndpointAddTransaction(s, logger)).Methods(http.MethodPost)
	r.Handle("/bank/{wallet_id}/commit", ports.NewEndpointApplyTransaction(s, logger)).Methods(http.MethodPost)
	r.Handle("/bank/{wallet_id}", ports.NewEndpointGetTransactions(s, logger)).Methods(http.MethodGet)

	// Register other routes as needed
	return r
}
