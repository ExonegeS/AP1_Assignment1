package ports

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ExonegeS/AP1_Assignment1/internal/library/service"
	"github.com/gorilla/mux"
)

func NewEndpointDeleteBook(s service.Service, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			bookID = mux.Vars(r)["book_id"]
		)

		var input = service.DeleteBookInput{
			BookID: bookID,
		}

		response, err := s.DeleteBook(input)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to delete book: %v", err), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		logger.Error("Book recieved successfully", "input", input)
	}
}
