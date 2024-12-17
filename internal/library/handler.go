package library

import (
	"log/slog"
	"net/http"

	"github.com/ExonegeS/AP1_Assignment1/internal/library/ports"
	"github.com/ExonegeS/AP1_Assignment1/internal/library/service"
	"github.com/gorilla/mux"
)

func NewHandler(s service.Service, logger *slog.Logger) http.Handler {
	r := mux.NewRouter()
	r.Handle("/library/book/{book_id}/get", ports.NewEndpointGetBook(s, logger)).Methods(http.MethodGet)
	r.Handle("/library/book/{book_id}/borrow", ports.NewEndpointBorrowBook(s, logger)).Methods(http.MethodPost)
	r.Handle("/library/book/{book_id}/return", ports.NewEndpointReturnBook(s, logger)).Methods(http.MethodPost)
	r.Handle("/library/book/list", ports.NewEndpointGetBooks(s, logger)).Methods(http.MethodGet)
	r.Handle("/library/book/add", ports.NewEndpointAddBook(s, logger)).Methods(http.MethodPost)

	// Register other routes as needed
	return r
}
