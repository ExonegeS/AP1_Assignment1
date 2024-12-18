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
	r.Handle("/library/books/add", ports.NewEndpointAddBook(s, logger)).Methods(http.MethodPost)
	r.Handle("/library/books/list", ports.NewEndpointGetBooks(s, logger)).Methods(http.MethodGet)
	r.Handle("/library/books/{book_id}", ports.NewEndpointGetBook(s, logger)).Methods(http.MethodGet)
	r.Handle("/library/books/{book_id}", ports.NewEndpointPutBook(s, logger)).Methods(http.MethodPut)
	r.Handle("/library/books/{book_id}", ports.NewEndpointDeleteBook(s, logger)).Methods(http.MethodDelete)
	r.Handle("/library/books/{book_id}/borrow", ports.NewEndpointBorrowBook(s, logger)).Methods(http.MethodPost)
	r.Handle("/library/books/{book_id}/return", ports.NewEndpointReturnBook(s, logger)).Methods(http.MethodPost)

	// Register other routes as needed
	return r
}
