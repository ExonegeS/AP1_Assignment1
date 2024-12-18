package ports

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ExonegeS/AP1_Assignment1/internal/library/service"
	"github.com/gorilla/mux"
)

func NewEndpointPutBook(s service.Service, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			bookID = mux.Vars(r)["book_id"]
		)
		var input = service.PutBookInput{
			BookID: bookID,
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			logger.Error("Invalid input", "input", input)
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		response, err := s.PutBook(input)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to update book: %v", err), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		logger.Error("Book recieved successfully", "input", input)
	}
}
